'use client';

interface InputProps {
  label: string;
  value: string | number;
  onChange: (value: string) => void;
  type?: 'text' | 'number';
  placeholder?: string;
  prefix?: string;
  suffix?: string;
  error?: string;
  min?: number;
  max?: number;
  step?: number;
}

export function Input({
  label,
  value,
  onChange,
  type = 'text',
  placeholder,
  prefix,
  suffix,
  error,
  min,
  max,
  step,
}: InputProps) {
  return (
    <div className="w-full">
      <label className="block text-sm text-[#a1a1aa] mb-2">{label}</label>
      <div className="relative flex items-center">
        {prefix && (
          <span className="absolute left-3 text-[#71717a] pointer-events-none">
            {prefix}
          </span>
        )}
        <input
          type={type}
          value={value}
          onChange={(e) => onChange(e.target.value)}
          placeholder={placeholder}
          min={min}
          max={max}
          step={step}
          className={`
            w-full bg-[#1c1c1f] border rounded-lg p-3 text-white
            placeholder:text-[#71717a]
            focus:outline-none focus:border-[#8b5cf6] focus:ring-1 focus:ring-[#8b5cf6]
            transition-all duration-200
            ${prefix ? 'pl-8' : ''}
            ${suffix ? 'pr-12' : ''}
            ${error ? 'border-red-500 focus:border-red-500 focus:ring-red-500' : 'border-[#3f3f46]'}
          `}
        />
        {suffix && (
          <span className="absolute right-3 text-[#71717a] pointer-events-none text-sm">
            {suffix}
          </span>
        )}
      </div>
      {error && (
        <p className="mt-1 text-xs text-red-400">{error}</p>
      )}
    </div>
  );
}