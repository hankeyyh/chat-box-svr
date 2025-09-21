"use client";

import { useState } from 'react';
import Toast from '../components/ui/Toast';
import Sidebar from '../components/layout/Sidebar';
import TopToolbar from '../components/layout/TopToolbar';
import ChatMessage from '../components/chat/ChatMessage';
import InteractionButtons from '../components/chat/InteractionButtons';
import InputArea from '../components/chat/InputArea';

export default function Home() {
  const [showToast, setShowToast] = useState(false);
  const [toastMessage, setToastMessage] = useState('');
  const [inputValue, setInputValue] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [streamingText, setStreamingText] = useState('');
  const [userQuestion, setUserQuestion] = useState('');

  const handleButtonClick = async () => {
    if (!inputValue.trim()) {
      setToastMessage('请输入内容');
      setShowToast(true);
      setTimeout(() => {
        setShowToast(false);
      }, 3000);
      return;
    }

    setIsLoading(true);
    setStreamingText('');
    setUserQuestion(inputValue);

    try {
      // 使用 Server-Sent Events 接收流式数据
      const response = await fetch('/api/chat', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ text: inputValue }),
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const reader = response.body?.getReader();
      const decoder = new TextDecoder();

      if (!reader) {
        throw new Error('无法读取响应流');
      }

      while (true) {
        const { done, value } = await reader.read();

        if (done) break;

        const chunk = decoder.decode(value);
        const lines = chunk.split('\n');

        for (const line of lines) {
          if (!line.startsWith('data: ')) {
            continue;
          }
          try {
            const data = JSON.parse(line.slice(6));

            if (data.type === 'connected') {
              console.log('SSE 连接已建立');
            } else if (data.type === 'data') {
              const response = data.response;

              // 显示每次接收到的数据
              if (response.data && response.data.text) {
                setStreamingText(prev => prev + response.data.text);
              }
            } else if (data.type === 'end') {
              console.log('流式响应结束');
              setToastMessage('响应完成');
              setShowToast(true);
              setTimeout(() => {
                setShowToast(false);
              }, 2000);
            } else if (data.type === 'error') {
              throw new Error(data.error);
            }
          } catch (parseError) {
            console.error('解析 SSE 数据失败:', parseError);
          }
        }
      }

    } catch (error) {
      console.error('发送消息失败:', error);
      setToastMessage('发送消息失败，请检查网络连接');
      setShowToast(true);
      setTimeout(() => {
        setShowToast(false);
      }, 3000);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-white flex">
      <Toast show={showToast} message={toastMessage} />

      <Sidebar />

      <div className="flex-1 flex flex-col">
        <TopToolbar />

        <ChatMessage
          userQuestion={userQuestion}
          streamingText={streamingText}
          isLoading={isLoading}
        />

        <InteractionButtons show={!!streamingText && !isLoading} />

        <InputArea
          inputValue={inputValue}
          isLoading={isLoading}
          onInputChange={setInputValue}
          onSend={handleButtonClick}
        />
      </div>
    </div>
  );
}
