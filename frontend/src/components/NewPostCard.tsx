interface NewPostCardProps {
    onClick: () => void;
}

export function NewPostCard({ onClick }: NewPostCardProps) {
    return (
        <button
            onClick={onClick}
            className="button-block sqs-block-button-element min-w-[280px] max-w-[300px] mt-[10px] h-[340px] flex flex-col items-center justify-center rounded-lg border-[1px] border-dashed border-[#D5D7DA] text-[#717680] hover:bg-gray-100"
        >
            <div className="flex flex-col items-center gap-0">
                <svg
                    width="45"
                    height="45"
                    viewBox="0 0 30 30"
                    fill="none"
                    stroke="currentColor"
                    className="text-[#717680]"
                >
                    <circle cx="12" cy="12" r="7" strokeWidth="1.5"/>
                    <path d="M12 9v6m-3-3h6" strokeWidth="1.5" strokeLinecap="round"/>
                </svg>
                <span className="text-[16px] font-inter font-normal">New Post</span>
            </div>
        </button>
    );
}
