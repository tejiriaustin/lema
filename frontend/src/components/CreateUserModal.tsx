import {Modal} from "./Modals.tsx";
import React, {useState} from "react";
import {InputBox} from "./InputBox.tsx";
import {Button} from "./Button.tsx";
import {Address} from "../types";

interface CreateUserModalProps {
    isOpen: boolean;
    onClose: () => void;
    onSubmit: (data: { full_name: string; email: string; address: Address}) => void;
}

export function CreateUserModal({isOpen, onSubmit, onClose}: CreateUserModalProps) {
    const [formData, setFormData] = useState({
        full_name: '',
        email: '',
        address: {
            street: '',
            city: '',
            state: '',
            zipcode: ''
        }
    });

    const handleSubmit = (e: React.FormEvent) => {
      e.preventDefault();
      onSubmit(formData);
    };

    const updateAddress = (field: string, value: string) => {
        setFormData(prev => ({
            ...prev,
            address: {
                ...prev.address,
                [field]: value
            }
        }));
    };

    return (
        <Modal isOpen={isOpen} onClose={onClose} title="Create User">
            <form onSubmit={handleSubmit} className="space-y-4">
                <InputBox
                    label="Full Name"
                    placeholder="David Beckham"
                    value={formData.full_name}
                    onChange={(value) => setFormData({ ...formData, full_name: value })}
                />
                <InputBox
                    label="Email"
                    placeholder="beckham@gmail.com"
                    value={formData.email}
                    onChange={(value) => setFormData({ ...formData, email: value })}
                    type="email"
                />
                <InputBox
                    label="Street Address"
                    value={formData.address.street}
                    onChange={(value) => updateAddress('street', value)}
                    placeholder="123 Main St"
                />
                <div className="grid grid-cols-2 gap-4">
                    <InputBox
                        label="City"
                        value={formData.address.city}
                        onChange={(value) => updateAddress('city', value)}
                        placeholder="Boston"
                    />
                    <InputBox
                        label="State"
                        value={formData.address.state}
                        onChange={(value) => updateAddress('state', value)}
                        placeholder="MA"
                    />
                </div>
                <InputBox
                    label="ZIP Code"
                    value={formData.address.zipcode}
                    onChange={(value) => updateAddress('zipcode', value)}
                    placeholder="02108"
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
                        disabled={!formData.full_name || !formData.email}
                    />
                </div>
            </form>
        </Modal>
    )
}