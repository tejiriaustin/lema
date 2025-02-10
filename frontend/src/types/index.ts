export interface User {
    id: string;
    fullName: string;
    email: string;
    address: string | Address;
}

export interface Post {
    id: string;
    userId: string;
    title: string;
    body: string;
}

export interface Address {
    street: string;
    city: string;
    state: string;
    zipcode: string;
}

export interface PaginatedResponse<T> {
    data: T[];
    total: number;
    page: number;
    pageSize: number;
}