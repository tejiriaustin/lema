interface PostCardProps {
    title: string;
    content: string;
    onDelete?: () => void;
}

function truncateText(text: string, limit: number): string {
    if (text.length <= limit) return text;
    return text.slice(0, limit).trim() + '...';
}

export function PostCard({ title, content, onDelete }: PostCardProps) {
    const truncatedTitle = truncateText(title, 50);

    return (
        <div className="rounded-lg border border-gray-300 p-6 relative shadow-lg w-[310px] h-[340px]">
            <h2
                className="text-[20px] text-[#535862] font-inter mb-4 overflow-hidden text-ellipsis"
                title={title}
            >
                {truncatedTitle}
            </h2>
            <p className="text-[16px] font-[300] text-[#535862] font-inter clamp-text">{content}</p>

            {onDelete && (
                <button
                    onClick={onDelete}
                    className="absolute top-3 right-3 text-red-500 hover:text-red-600"
                >
                    <svg
                        width="14"
                        height="14"
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