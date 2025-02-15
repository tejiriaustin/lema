import {Modal} from "./Modals.tsx";
import {Button} from "./Button.tsx";

interface DeleteConfirmationModalProps {
    isOpen: boolean;
    onClose: () => void;
    onConfirm: () => void;
    isLoading?: boolean;
}

export function DeleteConfirmationModal({ isOpen, onClose, onConfirm, isLoading }: DeleteConfirmationModalProps) {
    return (
        <Modal isOpen={isOpen} onClose={onClose} title="Delete Post">
            <div className="space-y-4">
                <p className="text-gray-600">
                    Are you sure you want to delete this post? This action cannot be undone.
                </p>
                <div className="flex justify-end gap-3 mt-6">
                    <Button
                        title="Cancel"
                        onClick={onClose}
                    />
                    <Button
                        title="Delete"
                        variant="primary"
                        onClick={() => {
                            onConfirm();
                            onClose();
                        }}
                        isLoading={isLoading}
                    />
                </div>
            </div>
        </Modal>
    );
}