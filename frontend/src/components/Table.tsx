import { ReactNode, useState } from "react";
import { Loader } from "./Loader";
import { ArrowRight, ArrowLeft } from "lucide-react";

interface Column<T> {
    header: string;
    key: keyof T;
    width?: string;
    render?: (value: T[keyof T], item: T) => ReactNode;
}

interface TableProps<T> {
    data: T[];
    columns: Column<T>[];
    onRowClick?: (item: T) => void;
    isLoading?: boolean;
    rowsPerPage?: number;
    currentPage: number
    onPageChange: (page: number) => void
    totalPages: number
}

export function Table<T>({data, columns, onRowClick, isLoading, rowsPerPage = 4,}: TableProps<T>) {
    const [currentPage, setCurrentPage] = useState(1);
    const totalPages = Math.ceil(data.length / rowsPerPage);

    const paginatedData = data.slice(
        (currentPage - 1) * rowsPerPage,
        currentPage * rowsPerPage
    );

    const goToPage = (page: number) => {
        if (page >= 1 && page <= totalPages) {
            setCurrentPage(page);
        }
    };

    const renderPaginationNumbers = () => {
        const pages = [];
        const showEllipsis = totalPages > 7;

        if (showEllipsis) {
            if (currentPage <= 4) {
                for (let i = 1; i <= 5; i++) pages.push(i);
                pages.push('...');
                pages.push(totalPages);
            } else if (currentPage >= totalPages - 3) {
                pages.push(1);
                pages.push('...');
                for (let i = totalPages - 4; i <= totalPages; i++) pages.push(i);
            } else {
                pages.push(1);
                pages.push('...');
                for (let i = currentPage - 1; i <= currentPage + 1; i++) pages.push(i);
                pages.push('...');
                pages.push(totalPages);
            }
        } else {
            for (let i = 1; i <= totalPages; i++) pages.push(i);
        }

        return pages;
    };

    return (
        <div className="flex flex-col">
            <div className="overflow-x-auto">
                <div className="rounded-lg shadow-sm overflow-hidden border border-gray-200 font-inter min-w-[640px]">
                    <table className="w-full border-collapse">
                        <thead>
                        <tr>
                            {columns.map((column) => (
                                <th
                                    key={String(column.key)}
                                    className="text-left text-[14px] font-medium text-gray-700 py-5 px-6"
                                    style={{width: column.width}}
                                >
                                    {column.header}
                                </th>
                            ))}
                        </tr>
                        </thead>
                        <tbody>
                        {isLoading ? (
                            <tr>
                                <td colSpan={columns.length} className="h-64">
                                    <div className="flex items-center justify-center h-full">
                                        <Loader/>
                                    </div>
                                </td>
                            </tr>
                        ) : paginatedData.length > 0 ? (
                            paginatedData.map((item, index) => (
                                <tr
                                    key={index}
                                    onClick={() => onRowClick?.(item)}
                                    className="cursor-pointer border-b last:border-none hover:bg-gray-50"
                                >
                                    {columns.map((column) => (
                                        <td key={String(column.key)}
                                            className={`text-[16px] py-7 px-6 tracking-wide ${
                                                column.key === 'address' ? 'truncate max-w-[1px]' : ''
                                            } ${
                                                column.key === 'fullName'
                                                    ? 'font-[475] text-[#535862]'
                                                    : 'font-[425] text-gray-500'
                                            }`}>
                                            {column.render
                                                ? column.render(item[column.key], item)
                                                : String(item[column.key] || "")}
                                        </td>
                                    ))}
                                </tr>
                            ))
                        ) : (
                            <tr>
                                <td colSpan={columns.length} className="text-center text-gray-500 py-4">
                                    No data available.
                                </td>
                            </tr>
                        )}
                        </tbody>
                    </table>
                </div>
            </div>

            <div className="flex justify-end mt-4 px-6 py-4 bg-white">
                <div className="flex items-center gap-2">
                    <button
                        className={`flex items-center gap-3 text-[16px] ${
                            currentPage === 1 ? "text-gray-400 cursor-not-allowed" : "text-gray-800 hover:text-gray-900"
                        }`}
                        onClick={() => goToPage(currentPage - 1)}
                        disabled={currentPage === 1}
                    >
                        <ArrowLeft size={18}/>
                        Previous
                    </button>

                    <div className="flex gap-2">
                        {renderPaginationNumbers().map((page, index) => (
                            <button
                                key={index}
                                className={`w-8 h-8 flex items-center justify-center text-sm rounded-lg ${
                                    page === '...'
                                        ? 'cursor-default'
                                        : page === currentPage
                                            ? "bg-purple-200 text-purple-800 font-semibold"
                                            : "text-gray-700 hover:bg-gray-100"
                                }`}
                                onClick={() => typeof page === 'number' && goToPage(page)}
                                disabled={page === '...'}
                            >
                                {page}
                            </button>
                        ))}
                    </div>

                    <button
                        className={`flex items-center gap-3 text-[16px] ${
                            currentPage === totalPages ? "text-gray-400 cursor-not-allowed" : "text-gray-700 hover:text-gray-900"
                        }`}
                        onClick={() => goToPage(currentPage + 1)}
                        disabled={currentPage === totalPages}
                    >
                        Next
                        <ArrowRight size={18}/>
                    </button>
                </div>
            </div>
        </div>
    );
}
