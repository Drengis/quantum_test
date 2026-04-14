'use client';

interface SliderProps {
  label: string;
  value: number;
  onChange: (value: number) => void;
  min: number;
  max: number;
  step?: number;
  suffix?: string;
  formatValue?: (value: number) => string;
}

export function Slider({
  label,
  value,
  onChange,
  min,
  max,
  step = 1,
  suffix,
  formatValue,
}: SliderProps) {
  const displayValue = formatValue ? formatValue(value) : value.toString();

  return (
    <div className="w-full">
      <div className="flex justify-between items-center mb-2">
        <label className="text-sm text-[#a1a1aa]">{label}</label>
        <span className="text-sm font-medium text-white">
          {displayValue}{suffix && ` ${suffix}`}
        </span>
      </div>
      <div className="relative">
        <input
          type="range"
          value={value}
          onChange={(e) => onChange(Number(e.target.value))}
          min={min}
          max={max}
          step={step}
          className="
            w-full h-2 bg-[#252529] rounded-lg appearance-none cursor-pointer
            [&::-webkit-slider-thumb]:appearance-none
            [&::-webkit-slider-thumb]:w-5
            [&::-webkit-slider-thumb]:h-5
            [&::-webkit-slider-thumb]:bg-[#8b5cf6]
            [&::-webkit-slider-thumb]:rounded-full
            [&::-webkit-slider-thumb]:shadow-lg
            [&::-webkit-slider-thumb]:cursor-pointer
            [&::-webkit-slider-thumb]:transition-transform
            [&::-webkit-slider-thumb]:active:scale-110
            [&::-moz-range-thumb]:w-5
            [&::-moz-range-thumb]:h-5
            [&::-moz-range-thumb]:bg-[#8b5cf6]
            [&::-moz-range-thumb]:rounded-full
            [&::-moz-range-thumb]:border-0
            [&::-moz-range-thumb]:cursor-pointer
          "
        />
        <div
          className="absolute top-1/2 -translate-y-1/2 h-2 bg-[#8b5cf6] rounded-l-lg pointer-events-none"
          style={{
            left: `${((value - min) / (max - min)) * 100}%`,
            width: `${((value - min) / (max - min)) * 100}%`,
          }}
        />
      </div>
      <div className="flex justify-between mt-1">
        <span className="text-xs text-[#71717a]">{min}{suffix}</span>
        <span className="text-xs text-[#71717a]">{max}{suffix}</span>
      </div>
    </div>
  );
}