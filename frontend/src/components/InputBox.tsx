interface InputProps {
    label: string;
    textLimit?: number;
    type?: string;
    placeholder: string;
    value: string;
    onChange: (value: string) => void;
    onExceeded?: (isExceeded: boolean) => void;
    multiline?: boolean;
    error?: string;
}

export function InputBox({label, textLimit, type = "text", placeholder, value, onChange, onExceeded, multiline = false,  error}: InputProps) {
    const baseStyles = "w-full px-4 py-3 rounded-lg border border-gray-200 focus:outline-none focus:border-gray-300 text-gray-900 placeholder:text-gray-400 font-inter text-[14px]";
    const isExceeded = textLimit && value.length > textLimit;

    const handleChange = (newValue: string) => {
        onChange(newValue);
        if (onExceeded) {
            onExceeded(textLimit ? newValue.length > textLimit : false);
        }
    };

    return (
        <div className="space-y-2">
            <div className="flex justify-between items-center">
                <label className="block text-[14px] font-inter text-gray-700">
                    {label}
                </label>
                {textLimit && (
                    <span className={`text-[12px] font-inter ${isExceeded ? 'text-red-500' : 'text-gray-500'}`}>
                        {value.length}/{textLimit}
                    </span>
                )}
            </div>

            {multiline ? (
                <textarea
                    value={value}
                    onChange={(e) => handleChange(e.target.value)}
                    placeholder={placeholder}
                    className={`${baseStyles} min-h-[200px] resize-none ${isExceeded ? 'border-red-500' : ''}`}
                />
            ) : (
                <input
                    type={type}
                    value={value}
                    onChange={(e) => handleChange(e.target.value)}
                    placeholder={placeholder}
                    className={`${baseStyles} ${isExceeded ? 'border-red-500' : ''}`}
                />
            )}

            {error && (
                <p className="text-red-500 text-[12px]">{error}</p>
            )}

            {isExceeded && (
                <p className="text-red-500 text-[12px]">
                    Text exceeds maximum limit of {textLimit} characters
                </p>
            )}
        </div>
    );
}