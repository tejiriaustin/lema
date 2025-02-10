interface ButtonProps {
    title: string;
    variant?: 'primary' | 'secondary';
    isLoading?: boolean;
    onClick?: () => void;
    type?: 'button' | 'submit';
    disabled?: boolean;
    className?: string;
}

export function Button({title, variant = 'secondary', isLoading = false, onClick, type = 'button', disabled = false, className}: ButtonProps) {
    const baseStyles = "px-6 py-2 rounded-md text-[14px] font-inter transition-colors";
    const variants = {
        primary: "bg-[#334155] text-white hover:bg-gray-500",
        secondary: "bg-white text-[#1F2937] border border-gray-200 hover:bg-gray-100"
    };

    return (
        <button
            type={type}
            onClick={onClick}
            disabled={disabled || isLoading}
            className={`${baseStyles} ${variants[variant]} ${className}`}
        >
            {isLoading ? (
                <div className="flex items-center justify-center">
                    <div className="w-5 h-5 border-t-2 border-b-2 border-current rounded-full animate-spin mr-2"></div>
                    {title}
                </div>
            ) : (
                title
            )}
        </button>
    );
}