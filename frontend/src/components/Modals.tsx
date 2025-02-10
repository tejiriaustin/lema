import React, { useState, useEffect } from 'react';

interface ModalProps {
    isOpen: boolean;
    onClose: () => void;
    children: React.ReactNode;
    title: string;
}

export function Modal({ isOpen, onClose, children, title }: ModalProps) {
    const [isAnimatingOut, setIsAnimatingOut] = useState(false);
    const [shouldRender, setShouldRender] = useState(false);

    useEffect(() => {
        if (isOpen) {
            setShouldRender(true);
            setIsAnimatingOut(false);
        } else if (shouldRender) {
            setIsAnimatingOut(true);
            const timer = setTimeout(() => {
                setShouldRender(false);
            }, 200);
            return () => clearTimeout(timer);
        }
    }, [isOpen]);

    if (!shouldRender) return null;

    const handleBackdropClick = (e: React.MouseEvent) => {
        if (e.target === e.currentTarget) {
            handleClose();
        }
    };

    const handleClose = () => {
        setIsAnimatingOut(true);
        setTimeout(() => {
            onClose();
        }, 200);
    };

    return (
        <div
            className={`fixed inset-0 flex items-center justify-center z-50 
                transition-opacity duration-200 animate-[fadeIn_200ms_ease-out]
                ${isAnimatingOut ? 'opacity-0' : 'opacity-100 bg-black bg-opacity-30'}`}
            onClick={handleBackdropClick}
        >
            <div
                className={`bg-white rounded-lg w-full max-w-[640px] p-6 
                    transition-all duration-200 animate-[slideUp_200ms_ease-out]
                    ${isAnimatingOut ? 'opacity-0 translate-y-4' : 'opacity-100 translate-y-0'}`}
                onClick={e => e.stopPropagation()}
            >
                <div className="flex justify-between items-center mb-6">
                    <h2 className="text-xl font-medium">{title}</h2>
                    <button
                        onClick={handleClose}
                        className="text-gray-400 hover:text-gray-600"
                    >
                        <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor">
                            <path
                                d="M18 6L6 18M6 6l12 12"
                                strokeWidth="2"
                                strokeLinecap="round"
                            />
                        </svg>
                    </button>
                </div>
                {children}
            </div>
        </div>
    );
}