'use client';

interface SelectOption {
    value: string;
    label: string;
}

interface SelectProps {
    label?: string;
    value: string;
    onChange: (value: string) => void;
    options: SelectOption[];
    placeholder?: string;
    className?: string;
}

export function Select({ label, value, onChange, options, placeholder, className = '' }: SelectProps) {
    return (
        <div className={`w-full ${className}`}>
            {label && (
                <label className="block text-sm text-[#a1a1aa] mb-2">
                    {label}
                </label>
            )}
            <select
                value={value}
                onChange={(e) => onChange(e.target.value)}
                className="w-full bg-[#1c1c1f] border border-[#3f3f46] rounded-xl p-4 text-white transition-all hover:border-[#8b5cf6] focus:outline-none focus:border-[#8b5cf6] appearance-none cursor-pointer"
            >
                {placeholder && (
                    <option value="" disabled>
                        {placeholder}
                    </option>
                )}
                {options.map((option) => (
                    <option key={option.value} value={option.value}>
                        {option.label}
                    </option>
                ))}
            </select>
        </div>
    );
}