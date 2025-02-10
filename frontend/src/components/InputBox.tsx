interface InputProps {
    label: string;
    type?: string;
    placeholder: string;
    value: string;
    onChange: (value: string) => void;
    multiline?: boolean;
    error?: string;
}

export function InputBox({label, type="text",placeholder,value,onChange,multiline = false,error}: InputProps) {
    const baseStyles = "w-full px-4 py-3 rounded-lg border border-gray-200 focus:outline-none focus:border-gray-300 text-gray-900 placeholder:text-gray-400 font-inter text-[14px]";

    return (
        <div className="space-y-2">
            <label className="block text-[14px] font-inter text-gray-700">
                {label}
            </label>

            {multiline ? (
                <textarea
                    value={value}
                    onChange={(e) => onChange(e.target.value)}
                    placeholder={placeholder}
                    className={`${baseStyles} min-h-[200px] resize-none`}
                />
            ) : (
                <input
                    type={type}
                    value={value}
                    onChange={(e) => onChange(e.target.value)}
                    placeholder={placeholder}
                    className={baseStyles}
                />
            )}

            {error && (
                <p className="text-red-500 text-[12px]">{error}</p>
            )}
        </div>
    );
}