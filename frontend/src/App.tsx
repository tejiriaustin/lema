import './App.css';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { ReactQueryDevtools } from '@tanstack/react-query-devtools';
import { UsersPage } from './pages/UsersPage'
import { UserPostsPage } from './pages/UserPostsPage'
import {TestPage} from "./pages/test.tsx";

const queryClient = new QueryClient({
        defaultOptions: {
            queries: {
                refetchOnWindowFocus: false,
                refetchOnMount: false,
                staleTime: 30000,
                retry: false,
                refetchOnReconnect: false,
            },
        },
    }
);

function App() {
    return (
        <QueryClientProvider client={queryClient}>
            <BrowserRouter>
                <Routes>
                    <Route path="/test" element={<TestPage />} />
                    <Route path="/" element={<UsersPage />} />
                    <Route path="/users/:userId/posts" element={<UserPostsPage />} />
                </Routes>
            </BrowserRouter>
            <ReactQueryDevtools />
        </QueryClientProvider>
    )
}

export default App