interface PostCardProps {
    title: string;
    content: string;
    onDelete?: () => void;
}

export function PostCard({ title, content, onDelete }: PostCardProps) {
    return (
        <div className="rounded-lg mt-[10px] border border-gray-200 p-6 relative shadow-sm min-w-[280px] max-w-[300px] h-[340px]">
            <h2 className="text-[20px] text-[#535862] font-inter mb-4">{title}</h2>
            <p className="text-[14px] text-[#535862] font-inter clamp-text">{content}</p>

            {onDelete && (
                <button
                    onClick={onDelete}
                    className="absolute top-3 right-4 text-red-500 hover:text-red-600"
                >
                    <svg
                        width="16"
                        height="16"
                        viewBox="0 0 16 16"
                        fill="none"
                        xmlns="http://www.w3.org/2000/svg"
                    >
                        <path
                            d="M2 4h12M5.333 4V2.667a1.333 1.333 0 011.334-1.334h2.666a1.333 1.333 0 011.334 1.334V4m2 0v9.333a1.333 1.333 0 01-1.334 1.334H4.667a1.333 1.333 0 01-1.334-1.334V4h9.334z"
                            stroke="currentColor"
                            strokeWidth="1.2"
                            strokeLinecap="round"
                            strokeLinejoin="round"
                        />
                        <path
                            d="M6.667 7v5M9.333 7v5"
                            stroke="currentColor"
                            strokeWidth="1.2"
                            strokeLinecap="round"
                            strokeLinejoin="round"
                        />
                    </svg>
                </button>
            )}
        </div>
    );
}
