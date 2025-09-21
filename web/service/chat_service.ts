import path from 'path';
import * as grpc from '@grpc/grpc-js';
import * as protoLoader from '@grpc/proto-loader';

const hostAddr = "127.0.0.1:3344";

// 使用绝对路径指向项目根目录的 proto 文件
var protoPath = path.join(process.cwd(), '../proto', 'chat.proto');

var packageDefinition = protoLoader.loadSync(
  protoPath,
  {keepCase: true,
   longs: String,
   enums: String,
   defaults: true,
   oneofs: true
  });
var chatProto = grpc.loadPackageDefinition(packageDefinition).chat;

// 流式聊天函数 - 支持回调函数实时接收数据
export function chatStream(text: string, onData: (response: any) => void): Promise<void> {
  return new Promise((resolve, reject) => {
    // @ts-ignore
    var client = new chatProto.ChatService(hostAddr,
      grpc.credentials.createInsecure());

    const call = client.chatStream({
      text: text,
    });

    call.on('data', function(response: any) {
      onData(response);
    });

    call.on('end', function() {
      resolve();
    });

    call.on('error', function(err: any) {
      reject(err);
    });
  });
}
