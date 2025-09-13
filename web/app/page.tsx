"use client";

export default function Home() {
  return (
    <div className="min-h-screen bg-white flex">
      {/* 左侧边栏 */}
      <div className="w-80 bg-white border-r border-gray-200 flex flex-col">
        {/* 顶部区域 */}
        <div className="p-4 border-b border-gray-100">
          <div className="flex items-center gap-3 mb-4">
            {/* Logo */}
            <div className="w-8 h-8 bg-gray-800 rounded-full flex items-center justify-center">
              <span className="text-white text-sm font-bold">G</span>
            </div>
            <div className="flex items-center gap-1">
              <span className="text-gray-800 font-medium">Grok-3</span>
              <svg className="w-4 h-4 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 9l-7 7-7-7" />
              </svg>
            </div>
          </div>
          {/* 新聊天按钮 */}
          <button className="w-full flex items-center gap-3 p-3 bg-gray-50 hover:bg-gray-100 rounded-lg transition-colors">
            <svg className="w-5 h-5 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
            </svg>
            <span className="text-gray-800">新聊天</span>
          </button>
        </div>

        {/* 中间空白区域 - 用于显示聊天历史 */}
        <div className="flex-1 p-4">
          {/* 这里可以添加聊天历史列表 */}
        </div>

        {/* 底部导航和用户信息 */}
        <div className="p-4 border-t border-gray-100 space-y-2">
          {/* 导航菜单 */}
          <nav className="space-y-2">
            <a href="#" className="flex items-center gap-3 p-3 text-gray-700 hover:bg-gray-50 rounded-lg transition-colors">
              <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />
              </svg>
              <span>回到首页</span>
            </a>
            <a href="#" className="flex items-center gap-3 p-3 text-gray-700 hover:bg-gray-50 rounded-lg transition-colors">
              <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
              </svg>
              <span>一键换车</span>
            </a>
            <a href="#" className="flex items-center gap-3 p-3 text-gray-700 hover:bg-gray-50 rounded-lg transition-colors">
              <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
              <span>常见问题</span>
            </a>
            <a href="#" className="flex items-center gap-3 p-3 text-gray-700 hover:bg-gray-50 rounded-lg transition-colors">
              <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 11a7 7 0 01-7 7m0 0a7 7 0 01-7-7m7 7v4m0 0H8m4 0h4m-4-8a3 3 0 01-3-3V5a3 3 0 116 0v6a3 3 0 01-3 3z" />
              </svg>
              <span>语音聊天</span>
            </a>
          </nav>

          {/* 会员选项 */}
          <a href="#" className="flex items-center gap-3 p-3 text-gray-700 hover:bg-gray-50 rounded-lg transition-colors">
            <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1" />
            </svg>
            <div className="flex flex-col">
              <span className="text-sm">次数不够?开通会员</span>
              <span className="text-xs text-gray-500">解锁高额使用次数及全模型</span>
            </div>
          </a>

          {/* 用户信息 */}
          <div className="flex items-center gap-3 p-3 text-gray-700">
            <div className="w-8 h-8 bg-blue-500 rounded-full flex items-center justify-center">
              <svg className="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
              </svg>
            </div>
            <span>用户20139</span>
          </div>
        </div>
      </div>

      {/* 主内容区域 */}
      <div className="flex-1 flex flex-col">
        {/* 顶部工具栏 */}
        <div className="flex items-center justify-end p-4 border-b border-gray-100">
          <div className="flex items-center gap-3">
            <button className="p-2 hover:bg-gray-100 rounded-lg transition-colors">
              <svg className="w-5 h-5 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z" />
              </svg>
            </button>
            <button className="w-8 h-8 bg-gray-800 hover:bg-gray-900 rounded-lg flex items-center justify-center transition-colors">
              <svg className="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
              </svg>
            </button>
          </div>
        </div>

        {/* 聊天内容区域 */}
        <div className="flex-1 p-6 overflow-y-auto">
          {/* 用户问题 */}
          <div className="mb-6">
            <div className="text-gray-800 text-lg leading-relaxed">
              截至2025.9.5号,市面上用于聊大,图片处理,视频处理的土流āi 模型有哪些
            </div>
          </div>

          {/* AI回复 */}
          <div className="mb-8">
            <div className="text-gray-700 leading-relaxed">
              <div className="mb-4">
                As of September 5, 2025, the mainstream AI large models for chat, image processing, and video processing include:
              </div>
              
              <div className="mb-4">
                <strong>Chat:</strong>
                <ul className="list-disc list-inside ml-4 mt-2 space-y-1">
                  <li><strong>Grok 3 (xAI)</strong> - Advanced conversational AI with real-time capabilities</li>
                  <li><strong>ChatGPT (OpenAI)</strong> - Leading conversational AI model with multimodal capabilities</li>
                  <li><strong>Claude (Anthropic)</strong> - AI assistant focused on helpfulness and safety</li>
                  <li><strong>Gemini (Google)</strong> - Multimodal AI model with strong reasoning capabilities</li>
                  <li><strong>LLaMA-based models (Meta AI and others)</strong> - Open-source foundation models</li>
                </ul>
              </div>

              <div className="mb-4">
                <strong>Image Processing:</strong>
                <ul className="list-disc list-inside ml-4 mt-2 space-y-1">
                  <li><strong>DALL-E 3 (OpenAI)</strong> - High-quality image generation from text prompts</li>
                  <li><strong>Stable Diffusion (Stability AI)</strong> - Open-source image generation model</li>
                  <li><strong>MidJourney</strong> - Popular AI art generation platform</li>
                  <li><strong>Imagen (Google)</strong> - Advanced image generation with high fidelity</li>
                  <li><strong>Flux.1 (Black Forest Labs)</strong> - State-of-the-art image synthesis model</li>
                </ul>
              </div>

              <div className="mb-4">
                <strong>Video Processing:</strong>
                <ul className="list-disc list-inside ml-4 mt-2 space-y-1">
                  <li><strong>Sora (OpenAI)</strong> - Revolutionary text-to-video generation model</li>
                  <li><strong>Runway Gen-2 (Runway)</strong> - AI-powered video generation and editing</li>
                  <li><strong>Pika.art</strong> - User-friendly AI video creation platform</li>
                  <li><strong>Kaiber</strong> - AI video generation and animation tools</li>
                  <li><strong>Luma AI (Dream Machine)</strong> - Advanced 3D and video generation</li>
                </ul>
              </div>

              <div className="text-gray-600">
                These models continue to dominate their respective domains, with ongoing developments in quality, efficiency, and accessibility. The landscape is rapidly evolving, and I'd be happy to search for the most current specific details if you'd like.
              </div>
            </div>
          </div>

          {/* 交互按钮 */}
          <div className="flex items-center gap-4 text-gray-500 mb-8">
            <button className="flex items-center gap-2 hover:text-gray-700 transition-colors">
              <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M14 10h4.764a2 2 0 011.789 2.894l-3.5 7A2 2 0 0115.263 21h-4.017c-.163 0-.326-.02-.485-.06L7 20m7-10V18m-7-8a2 2 0 012-2h4.5a2 2 0 012 2v1a2 2 0 01-2 2H9a2 2 0 01-2-2V10z" />
              </svg>
              <span>赞</span>
            </button>
            <button className="flex items-center gap-2 hover:text-gray-700 transition-colors">
              <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M10 14H5.236a2 2 0 01-1.789-2.894l3.5-7A2 2 0 019.737 3h4.018c.162 0 .326.02.485.06L15 4m-5 8V6m7 8a2 2 0 01-2 2h-4.5a2 2 0 01-2-2v-1a2 2 0 012-2H15a2 2 0 012 2v1z" />
              </svg>
              <span>踩</span>
            </button>
            <button className="flex items-center gap-2 hover:text-gray-700 transition-colors">
              <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15.536 8.464a5 5 0 010 7.072m2.828-9.9a9 9 0 010 12.728M5.586 15H4a1 1 0 01-1-1v-4a1 1 0 011-1h1.586l4.707-4.707C10.923 3.663 12 4.109 12 5v14c0 .891-1.077 1.337-1.707.707L5.586 15z" />
              </svg>
              <span>语音</span>
            </button>
            <button className="flex items-center gap-2 hover:text-gray-700 transition-colors">
              <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.367 2.684 3 3 0 00-5.367-2.684z" />
              </svg>
              <span>分享</span>
            </button>
            <button className="flex items-center gap-2 hover:text-gray-700 transition-colors">
              <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
              </svg>
            </button>
          </div>
        </div>

        {/* 输入区域 */}
        <div className="p-6 border-t border-gray-100">
          <div className="max-w-4xl mx-auto">
            <div className="relative">
              <input
                type="text"
                placeholder="+ ΣΣ"
                className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-400 focus:border-transparent"
              />
            </div>
            
            {/* 底部提示 */}
            <div className="flex items-center justify-between mt-4">
              <div className="text-sm text-gray-500">
                ChatGPT 也可能会犯错。请核查重要信息。
              </div>
              <button className="w-8 h-8 bg-gray-800 hover:bg-gray-900 rounded-full flex items-center justify-center transition-colors">
                <svg className="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8" />
                </svg>
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
