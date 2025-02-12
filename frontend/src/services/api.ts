import axios,  { AxiosError } from 'axios';
import {User} from "../types";
import { API_URL } from '../config/env';

const api = axios.create({
    baseURL: API_URL
})

interface ApiError {
    message: string;
    code?: string;
}

const handleError = (error: unknown) => {
    if (axios.isAxiosError(error)) {
        const axiosError = error as AxiosError<ApiError>;
        throw new Error(
            axiosError.response?.data?.message ||
            axiosError.message ||
            'An unexpected error occurred'
        );
    }
    throw error;
};

export const createUser = async (user: { full_name: string; email: string; address: { street: string; city: string; state: string; zipcode: string}}): Promise<User> => {
    try {
        const { data } = await api.post(`/users`, user);
        return data;
    } catch (error) {
        throw handleError(error);
    }
}

export const getUsers = async (page: number) => {
    try {
        const { data } = await api.get(`/users?pageNumber=${page}&pageSize=4`);
        return data;
    } catch (error) {
        throw handleError(error);
    }
}

export const getUserPosts = async (userId: string) => {
    try {
        const { data } = await api.get(`/posts?user_id=${userId}`);
        return data;
    } catch (error) {
        throw handleError(error)
    }
}

export const createPost = async (post: { title: string; body: string; user_id: string }) => {
    try {
        const { data } = await api.post('/posts', post)
        return data
    } catch (error) {
        throw handleError(error)
    }
}

export const deletePost = async (id: string) => {
    try {
        await api.delete(`/posts/${id}`)
    } catch (error) {
        throw handleError(error)
    }
}