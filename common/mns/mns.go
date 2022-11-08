package mns

import (
	"encoding/json"
	"errors"
	"fmt"
	ali_mns "github.com/aliyun/aliyun-mns-go-sdk"
	"ldm/common/config"
	"ldm/common/constant"
	"log"
	"sync"
)

var (
	Client ali_mns.MNSClient
)
var mnsOnce sync.Once
var mapTopicHandle sync.Map
var mapQueueHandle sync.Map

type MnsHandleFunc interface {
	Handle(data MnsMessageStruct) error
}
type MnsMessageStruct struct {
	//消息主题
	Subject string `json:"subject"`
	//消息内容
	Content string `json:"content"`
}

//初始化
func InitMns() {
	mnsOnce.Do(func() {
		mnsConfig := config.GlobalConfig.Mns
		Client = ali_mns.NewAliMNSClient(mnsConfig.Url, mnsConfig.AccessKeyId, mnsConfig.AccessKeySecret)
		//队列名列表
		queueNameArray := constant.MNS_QUENE_NAME_ARR
		//主题列表
		queueTopicArray := constant.MNS_TOPIC_NAME_ARR
		//创建队列
		queueManager := ali_mns.NewMNSQueueManager(Client)
		for _, queueName := range queueNameArray {
			queueName = getQueueRealName(queueName)
			//创建队列
			err := queueManager.CreateSimpleQueue(queueName)
			if err != nil && !ali_mns.ERR_MNS_QUEUE_ALREADY_EXIST_AND_HAVE_SAME_ATTR.IsEqual(err) {
				log.Fatal("创建队列失败，队列名：", queueName, "错误信息:", err.Error())
			}
			queueHandle := ali_mns.NewMNSQueue(queueName, Client)
			mapQueueHandle.Store(queueName, queueHandle)
		}
		//创建主题
		topicManager := ali_mns.NewMNSTopicManager(Client)
		for _, topicName := range queueTopicArray {
			topicName = getTopicRealName(topicName)
			err := topicManager.CreateSimpleTopic(topicName)
			if err != nil && !ali_mns.ERR_MNS_TOPIC_ALREADY_EXIST_AND_HAVE_SAME_ATTR.IsEqual(err) {
				log.Fatal("创建主题失败,主题名:", topicName, "错误信息:", err.Error())
			}
			topicHandle := ali_mns.NewMNSTopic(topicName, Client)
			//存储句柄
			mapTopicHandle.Store(topicName, topicHandle)
		}
	})
}

//订阅主题【指定队列接收主题消息】
func SubsribeTopic(queueName string, topicName string, subscriptionName string) error {
	queueName = getQueueRealName(queueName)
	topicName = getTopicRealName(topicName)
	val, ok := mapTopicHandle.Load(topicName)
	if !ok {
		return errors.New(fmt.Sprintf("topic not found,queueName:%s topicName:%s ", queueName, topicName))
	}
	topicHandle, ok := val.(ali_mns.AliMNSTopic)
	if !ok {
		return errors.New(fmt.Sprintf("topic not found!!,queueName:%s topicName:%s ", queueName, topicName))
	}
	//订阅主题endpoint设置为队列名称
	sub := ali_mns.MessageSubsribeRequest{
		Endpoint:            topicHandle.GenerateQueueEndpoint(queueName),
		NotifyContentFormat: ali_mns.SIMPLIFIED,
	}
	err := topicHandle.Subscribe(subscriptionName, sub)
	if err != nil && !ali_mns.ERR_MNS_SUBSCRIPTION_ALREADY_EXIST_AND_HAVE_SAME_ATTR.IsEqual(err) {
		return err
	}
	return nil
}

//发布消息到主题中
func PublishMnsTopicMessage(topicName string, subject string, data string) error {
	topicName = getTopicRealName(topicName)
	val, ok := mapTopicHandle.Load(topicName)
	if !ok {
		return errors.New(fmt.Sprintf("topic not found,topicName:%s ", topicName))
	}
	topicHandle, ok := val.(ali_mns.AliMNSTopic)
	if !ok {
		return errors.New(fmt.Sprintf("topic not found!!, topicName:%s ", topicName))
	}
	d := MnsMessageStruct{
		Subject: subject,
		Content: data,
	}
	bytesData, err := json.Marshal(d)
	if err != nil {
		return err
	}
	// 发布消息。
	msg := ali_mns.MessagePublishRequest{
		MessageBody: string(bytesData),
	}
	_, err = topicHandle.PublishMessage(msg)
	if err != nil {
		return err
	}
	return nil
}

//发布消息到队列中
func PublishMnsQueueMessage(queueName string, subject string, data ...string) error {
	queueName = getQueueRealName(queueName)
	val, ok := mapQueueHandle.Load(queueName)
	if !ok {
		return errors.New(fmt.Sprintf("queue not found,queueName:%s ", queueName))
	}
	queueHandle, ok := val.(ali_mns.AliMNSQueue)
	if !ok {
		return errors.New(fmt.Sprintf("queue not found!!, queueName:%s ", queueName))
	}
	sendSlice := make([]ali_mns.MessageSendRequest, 0, len(data))
	for _, v := range data {
		d := MnsMessageStruct{
			Subject: subject,
			Content: v,
		}
		bytesData, err := json.Marshal(d)
		if err != nil {
			return err
		}
		// 发布消息。
		sendSlice = append(sendSlice, ali_mns.MessageSendRequest{
			MessageBody: string(bytesData),
			Priority:    1, //优先级，默认1
		})
	}
	_, err := queueHandle.BatchSendMessage(sendSlice...)
	if err != nil {
		return err
	}
	return nil
}

//发布消息到[延时]队列中
func PublishMnsDelayQueueMessage(queueName string, subject string, delayTime int64, data ...string) error {
	queueName = getQueueRealName(queueName)
	val, ok := mapQueueHandle.Load(queueName)
	if !ok {
		return errors.New(fmt.Sprintf("queue not found,queueName:%s ", queueName))
	}
	queueHandle, ok := val.(ali_mns.AliMNSQueue)
	if !ok {
		return errors.New(fmt.Sprintf("queue not found!!, queueName:%s ", queueName))
	}
	sendSlice := make([]ali_mns.MessageSendRequest, 0, len(data))
	for _, v := range data {
		d := MnsMessageStruct{
			Subject: subject,
			Content: v,
		}
		bytesData, err := json.Marshal(d)
		if err != nil {
			return err
		}
		// 发布消息。
		sendSlice = append(sendSlice, ali_mns.MessageSendRequest{
			MessageBody:  string(bytesData),
			Priority:     1, //优先级，默认1
			DelaySeconds: delayTime,
		})
	}
	_, err := queueHandle.BatchSendMessage(sendSlice...)
	if err != nil {
		return err
	}
	return nil
}

//消费队列消息
func Consume(queueName string, handleFunc MnsHandleFunc, numOfmessage int32) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("捕获异常:", err)
		}
	}()
	queueName = getQueueRealName(queueName)
	if numOfmessage <= 0 || numOfmessage > 16 {
		numOfmessage = 10
	}
	respChan := make(chan ali_mns.BatchMessageReceiveResponse)
	errChan := make(chan error)
	defer close(respChan)
	defer close(errChan)
	queue := ali_mns.NewMNSQueue(queueName, Client)
	//拉取数据协程
	go func() {
		for {
			queue.BatchReceiveMessage(respChan, errChan, numOfmessage, 30)
		}
	}()
	for {
		select {
		case respList := <-respChan:
			dealwithMsg(queue, handleFunc, respList.Messages)
		case <-errChan:
			continue
		}
	}
}

//处理消息
func dealwithMsg(queue ali_mns.AliMNSQueue, handleFunc MnsHandleFunc, MessagesList []ali_mns.MessageReceiveResponse) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("捕获异常:", err)
		}
	}()
	//消费成功的消息
	successChan := make(chan string)
	wait := sync.WaitGroup{}
	for _, resp := range MessagesList {
		wait.Add(1)
		//处理业务
		go func(rs ali_mns.MessageReceiveResponse) {
			defer func() {
				if err := recover(); err != nil {
					log.Println("捕获异常:", err)
				}
			}()
			defer wait.Done()
			var info MnsMessageStruct
			if err := json.Unmarshal([]byte(rs.MessageBody), &info); err != nil {
				log.Println("json unmarshal err:", err)
				return
			}
			if err := handleFunc.Handle(info); err != nil {
				log.Println("mns消费数据失败,data:", rs.MessageBody, "err msg:", err.Error())
			} else {
				successChan <- rs.ReceiptHandle
				fmt.Println("mns消息消费成功,队列名：", queue.Name(), "队列数据:", rs.MessageBody)
			}
		}(resp)
	}
	go func() {
		wait.Wait()
		close(successChan)
	}()
	//最长16就给16吧
	successSlic := make([]string, 0, 16)
	for v := range successChan {
		successSlic = append(successSlic, v)
	}
	errRsp, err := queue.BatchDeleteMessage(successSlic...)
	if err != nil {
		log.Println("消费成功但是删除失败的消息,总数：", len(MessagesList), "失败：", len(errRsp.FailedMessages))
	} else {
		fmt.Println("批次删除已成功消费消息全部成功,数量为:", len(MessagesList))
	}
}
func getQueueRealName(queueName string) string {
	return fmt.Sprintf("%s-%s", getEnv(), queueName)
}
func getTopicRealName(topicName string) string {
	return fmt.Sprintf("%s-%s", getEnv(), topicName)
}
func getEnv() string {
	mnsConfig := config.GlobalConfig.Mns
	switch mnsConfig.Env {
	case "develop", "test", "production":
		return mnsConfig.Env
	default:
		return "develop"
	}
}
