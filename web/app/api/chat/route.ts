import { NextRequest } from 'next/server';
import { chatStream } from '../../../service/chat_service';

export async function POST(request: NextRequest) {
    try {
        const { text } = await request.json();

        if (!text || typeof text !== 'string') {
            return new Response(
                JSON.stringify({ error: '请提供有效的文本内容' }),
                {
                    status: 400,
                    headers: { 'Content-Type': 'application/json' }
                }
            );
        }

        // 创建 SSE 流
        const stream = new ReadableStream({
            start(controller) {
                // 发送初始连接确认
                const encoder = new TextEncoder();
                controller.enqueue(encoder.encode('data: {"type":"connected"}\n\n'));

                // 调用流式 gRPC 服务
                chatStream(text, (response: any) => {
                    // 将每个响应发送到客户端
                    const data = JSON.stringify({
                        type: 'data',
                        response: response
                    });
                    controller.enqueue(encoder.encode(`data: ${data}\n\n`));
                }).then(() => {
                    // 流结束
                    controller.enqueue(encoder.encode('data: {"type":"end"}\n\n'));
                    controller.close();
                }).catch((error: any) => {
                    // 错误处理
                    const errorData = JSON.stringify({
                        type: 'error',
                        error: error.message
                    });
                    controller.enqueue(encoder.encode(`data: ${errorData}\n\n`));
                    controller.close();
                });
            }
        });

        return new Response(stream, {
            headers: {
                'Content-Type': 'text/event-stream',
                'Cache-Control': 'no-cache',
                'Connection': 'keep-alive',
                'Access-Control-Allow-Origin': '*',
                'Access-Control-Allow-Headers': 'Cache-Control',
            },
        });
    } catch (error) {
        console.error('API 调用失败:', error);
        return new Response(
            JSON.stringify({ error: '服务器内部错误' }),
            {
                status: 500,
                headers: { 'Content-Type': 'application/json' }
            }
        );
    }
} 