interface ChatMessageProps {
  userQuestion?: string;
  streamingText?: string;
  isLoading?: boolean;
}

export default function ChatMessage({ userQuestion, streamingText, isLoading }: ChatMessageProps) {
  return (
    <div className="flex-1 p-6 overflow-y-auto">
      {/* 用户问题 */}
      {userQuestion && (
        <div className="mb-6">
          <div className="text-gray-800 text-lg leading-relaxed">
            {userQuestion}
          </div>
        </div>
      )}

      {/* AI回复 - 显示流式文本 */}
      {streamingText && (
        <div className="mb-8">
          <div className="text-gray-700 leading-relaxed">
            <div className="whitespace-pre-wrap">
              {streamingText}
              {isLoading && (
                <span className="inline-block w-2 h-5 bg-gray-400 ml-1 animate-pulse"></span>
              )}
            </div>
          </div>
        </div>
      )}
    </div>
  );
} 