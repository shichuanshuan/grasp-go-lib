package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
)

func main() {
	// ����һ������ȡ�����ܵ�������
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	// ģ��һ����ʱ�������� 3 ���ȡ��������

	//time.Sleep(3 * time.Second)

	// ִ�в��������������״̬
	if ctx != nil {
		select {
		case <-ctx.Done():
			fmt.Println("������ȡ��:", ctx.Err())
		default:
			fmt.Println("�����������")
			// ���������Ӿ���Ĳ�������
		}
	} else {
		fmt.Println("������Ϊ��")
	}
}