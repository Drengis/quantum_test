'use client';

import { useSearchParams, useRouter } from 'next/navigation';
import { useState, useEffect } from 'react';
import { api } from '@/shared/api';

interface MortgagePayment {
    totalPayment: string;
    repaymentOfMortgageBody: string;
    repaymentOfMortgageInterest: string;
    mortgageBalance: string;
}

interface MortgageResponse {
    id: string;
    status: string;
    monthlyPayment: string;
    totalPayment: string;
    totalOverpaymentAmount: string;
    possibleTaxDeduction: string;
    savingsDueMotherCapital: string;
    recommendedIncome: string;
    mortgagePaymentSchedule?: Record<string, Record<string, MortgagePayment>>;
}

const monthOrder: Record<string, number> = {
    'January': 1, 'February': 2, 'March': 3, 'April': 4,
    'May': 5, 'June': 6, 'July': 7, 'August': 8,
    'September': 9, 'October': 10, 'November': 11, 'December': 12
};

const monthNames: Record<string, string> = {
    'January': 'Январь', 'February': 'Февраль', 'March': 'Март',
    'April': 'Апрель', 'May': 'Май', 'June': 'Июнь',
    'July': 'Июль', 'August': 'Август', 'September': 'Сентябрь',
    'October': 'Октябрь', 'November': 'Ноябрь', 'December': 'Декабрь'
};

export default function SchedulePage() {
    const searchParams = useSearchParams();
    const router = useRouter();
    const id = searchParams?.get('id');
    const [data, setData] = useState<MortgageResponse | null>(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        if (id) {
            api.get<MortgageResponse>(`/mortgage-profiles/${id}`).then(setData).finally(() => setLoading(false));
        }
    }, [id]);

    if (loading) {
        return (
            <div className="flex items-center justify-center min-h-[50vh]">
                <div className="w-12 h-12 rounded-full bg-gradient-to-r from-[#8b5cf6] to-[#7c3aed] animate-pulse-subtle" />
            </div>
        );
    }

    if (!data?.mortgagePaymentSchedule) {
        return (
            <div className="p-4">
                <h1 className="text-2xl font-bold mb-4 animate-fade-in">График платежей</h1>
                <p className="text-[#a1a1aa] animate-fade-in delay-100">Нет данных о графике</p>
            </div>
        );
    }

    const schedule = data.mortgagePaymentSchedule;
    const years = Object.keys(schedule).sort();

    return (
        <div className="p-4 pb-20">
            <button
                onClick={() => router.push('/')}
                className="mb-4 flex items-center gap-2 text-[#a1a1aa] hover:text-white transition-colors animate-fade-in"
            >
                <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M10 19l-7-7m0 0l7-7m-7 7h18" />
                </svg>
                <span>На главную</span>
            </button>

            <h1 className="text-2xl font-bold mb-2 animate-fade-in-up delay-0">График платежей</h1>
            <p className="text-[#a1a1aa] mb-6 animate-fade-in-up delay-100">
                Ежемесячный платёж: <span className="text-[#8b5cf6] font-semibold">{data.monthlyPayment} ₽</span>
            </p>

            {years.map((year, yearIndex) => (
                <div key={year} className="mb-8 animate-fade-in-up" style={{ animationDelay: `${150 + yearIndex * 50}ms` }}>
                    <h2 className="text-lg font-semibold mb-3 text-[#8b5cf6] sticky top-0 bg-[#131316]/95 py-2 backdrop-blur-sm z-10">
                        {year}
                    </h2>

                    <div className="overflow-x-auto">
                        <table className="w-full text-xs">
                            <thead className="bg-[#1c1c1f] border border-[#3f3f46] rounded-xl">
                                <tr className="text-left">
                                    <th className="p-3 rounded-l-xl">Месяц</th>
                                    <th className="p-3">Платёж</th>
                                    <th className="p-3">Основной долг</th>
                                    <th className="p-3">Проценты</th>
                                    <th className="p-3 rounded-r-xl">Остаток</th>
                                </tr>
                            </thead>
                            <tbody>
                                {(() => {
                                    const months = Object.keys(schedule[year] || {});
                                    const sortedMonths = months.sort((a, b) => {
                                        return (monthOrder[a] || 0) - (monthOrder[b] || 0);
                                    });
                                    return sortedMonths.map((month, index) => {
                                        const payment = schedule[year]![month];
                                        const monthName = monthNames[month] || month;
                                        return (
                                            <tr
                                                key={month}
                                                className={`border-b border-[#3f3f46] hover:bg-[#1c1c1f]/50 transition-all duration-200 ${index % 2 === 0 ? 'bg-[#1c1c1f]/30' : ''
                                                    }`}
                                            >
                                                <td className="p-3 font-medium">{monthName}</td>
                                                <td className="p-3 text-[#8b5cf6] font-semibold">{payment.totalPayment} ₽</td>
                                                <td className="p-3 text-[#a1a1aa]">{payment.repaymentOfMortgageBody} ₽</td>
                                                <td className="p-3 text-[#a1a1aa]">{payment.repaymentOfMortgageInterest} ₽</td>
                                                <td className="p-3 text-[#a1a1aa]">{payment.mortgageBalance} ₽</td>
                                            </tr>
                                        );
                                    });
                                })()}
                            </tbody>
                        </table>
                    </div>
                </div>
            ))}
        </div>
    );
}