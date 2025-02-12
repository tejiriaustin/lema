import { useEffect } from 'react';
import { X } from 'lucide-react';

export type ToastType = 'success' | 'error' | 'info';

interface ToastProps {
    message: string;
    type: ToastType;
    onClose: () => void;
    duration?: number;
}

export function Toast({ message, type, onClose, duration = 5000 }: ToastProps) {
    useEffect(() => {
        const timer = setTimeout(() => {
            onClose();
        }, duration);

        return () => clearTimeout(timer);
    }, [duration, onClose]);

    const bgColor = {
        success: 'bg-green-50 border-green-500 text-green-800',
        error: 'bg-red-50 border-red-500 text-red-800',
        info: 'bg-blue-50 border-blue-500 text-blue-800'
    }[type];

    return (
        <div className={`fixed top-4 right-4 p-4 rounded-lg border ${bgColor} shadow-lg max-w-md animate-slide-in`}>
            <div className="flex items-center justify-between">
                <p className="text-sm font-medium">{message}</p>
                <button
                    onClick={onClose}
                    className="ml-4 text-gray-400 hover:text-gray-600"
                >
                    <X size={16} />
                </button>
            </div>
        </div>
    );
}