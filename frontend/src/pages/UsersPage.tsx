import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useUsers } from '../hooks/useUsers.tsx';
import { Table } from '../components/Table';
import { User, Address } from '../types';
import { Button } from "../components/Button.tsx";
import { CreateUserModal } from "../components/CreateUserModal.tsx";
import { useToast } from '../components/ToastProvider';

export function UsersPage() {
    const [page, setPage] = useState(1);
    const navigate = useNavigate();
    const [isModalOpen, setIsModalOpen] = useState(false);
    const { data, createUser, isLoading, mutationError } = useUsers(page);
    const { showToast } = useToast();

    const columns = [
        { header: "Full Name", key: "fullName" as const, width: "220px" },
        { header: "Email Address", key: "email" as const},
        {
            header: "Address",
            key: "address" as const,
            width: "392px",
            render: (value: Address | string) => {
                try {
                    const address = value as Address;
                    return `${address.street}, ${address.city}, ${address.state}, ${address.zipCode}`;
                } catch (error) {
                    console.log(error)
                    showToast('Error displaying address', 'error');
                    return 'Invalid address';
                }
            }
        }
    ]

    const handleRowClick = (user: User) => {
        try {
            navigate(`/users/${user.id}/posts`);
        } catch (error) {
            console.log(error)
            showToast('Error navigating to user posts', 'error');
        }
    }

    const handlePageChange = (newPage: number) => {
        try {
            setPage(newPage);
        } catch (error) {
            console.log(error)
            showToast('Error changing page', 'error');
        }
    }

    const handleSubmit = async (userData: {
        full_name: string;
        email: string;
        address: {
            street: string;
            city: string;
            state: string;
            zipcode: string
        }
    }) => {
        createUser(userData, {
            onSuccess: () => {
                showToast('User created successfully', 'success');
                setIsModalOpen(false);
            },
            onError: (error: Error) => {
                showToast(error.message || 'Failed to create user', 'error');
            }
        });
    }

    return (
        <div className="p-4 md:p-6 lg:p-8 max-w-[1100px] mx-auto mt-30">
            <div className="flex flex-row sm:flex-row justify-between items-start sm:items-center gap-4 mb-6">
                <h1 className="text-4xl md:text-5xl lg:text-7xl font-normal">Users</h1>
                <Button
                    title="Create User"
                    variant="primary"
                    onClick={() => setIsModalOpen(true)}
                />
            </div>

            <div className="rounded-lg shadow-sm -mx-4 sm:mx-0">
                <div className="overflow-x-auto">
                    <Table
                        isLoading={isLoading}
                        data={data?.data?.body.users || []}
                        columns={columns}
                        onRowClick={handleRowClick}
                        rowsPerPage={4}
                        currentPage={page}
                        onPageChange={handlePageChange}
                        totalPages={data?.data?.body.paginationData.total_pages || 1}
                    />
                </div>
            </div>

            <CreateUserModal
                onClose={() => setIsModalOpen(false)}
                isOpen={isModalOpen}
                onSubmit={handleSubmit}
                error={mutationError as Error}
            />
        </div>
    )
}