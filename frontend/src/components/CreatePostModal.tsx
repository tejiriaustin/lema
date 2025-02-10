import {useState} from "react";
import {Modal} from "./Modals.tsx";
import {InputBox} from "./InputBox.tsx";
import {Button} from "./Button.tsx";

interface CreatePostModalProps {
    isOpen: boolean;
    onClose: () => void;
    onSubmit: (data: { title: string; body: string }) => void;
}

export function CreatePostModal({ isOpen, onClose, onSubmit }: CreatePostModalProps) {
    const [title, setTitle] = useState('');
    const [content, setContent] = useState('');

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        onSubmit({ title, body: content });
        setTitle('');
        setContent('');
        onClose();
    };

    return (
        <Modal isOpen={isOpen} onClose={onClose} title="Create New Post">
            <form onSubmit={handleSubmit} className="space-y-4">
                <InputBox
                    label="Post title"
                    placeholder="Give your post a title"
                    value={title}
                    onChange={setTitle}
                />
                <InputBox
                    label="Post content"
                    placeholder="Write something mind-blowing"
                    value={content}
                    onChange={setContent}
                    multiline
                />
                <div className="flex justify-end gap-3 mt-6">
                    <Button
                        title="Cancel"
                        onClick={onClose}
                    />
                    <Button
                        title="Publish"
                        variant="primary"
                        type="submit"
                        disabled={!title || !content}
                    />
                </div>
            </form>
        </Modal>
    );
}