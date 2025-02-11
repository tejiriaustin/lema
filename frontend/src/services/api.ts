import axios from 'axios';
import {User} from "../types";
import { API_URL } from '../config/env';

const api = axios.create({
    baseURL: API_URL
})

export const createUser = async (user: { full_name: string; email: string; address: { street: string; city: string; state: string; zipcode: string}}): Promise<User> => {
    const { data } = await api.post(`/users`, user)
    return data
}

export const getUsers = async (page: number) => {
    const { data } = await api.get(`/users?pageNumber=${page}&pageSize=4`)
    return data
}

export const getUserPosts = async (userId: string) => {
    const { data } = await api.get(`/posts?user_id=${userId}`)
    return data
}

export const createPost = async (post: { title: string; body: string; user_id: string }) => {
    const { data } = await api.post('/posts', post)
    return data
}

export const deletePost = async (id: string) => {
    await api.delete(`/posts/${id}`)
}