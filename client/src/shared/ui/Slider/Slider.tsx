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

  const percentage = Math.min(100, Math.max(0, ((value - min) / (max - min)) * 100));

  return (
    <div className="w-full">
      <div className="flex justify-between items-center mb-2">
        <label className="text-sm text-[#a1a1aa]">{label}</label>
        <span className="text-sm font-medium text-white">
          {displayValue}{suffix && ` ${suffix}`}
        </span>
      </div>

      <div className="relative py-3">
        <div className="absolute w-full h-2 bg-[#252529] rounded-full overflow-hidden top-1/2 -translate-y-2">
          <div
            className="h-full bg-[#8b5cf6]"
            style={{
              width: `${percentage}%`,
              borderRadius: '9999px 0 0 9999px'
            }}
          />
        </div>

        <input
          type="range"
          value={value}
          onChange={(e) => onChange(Number(e.target.value))}
          min={min}
          max={max}
          step={step}
          className="
            relative
            w-full
            appearance-none
            bg-transparent
            cursor-pointer
            [&::-webkit-slider-thumb]:appearance-none
            [&::-webkit-slider-thumb]:relative
            [&::-webkit-slider-thumb]:z-20
            [&::-webkit-slider-thumb]:w-5
            [&::-webkit-slider-thumb]:h-5
            [&::-webkit-slider-thumb]:bg-[#8b5cf6]
            [&::-webkit-slider-thumb]:rounded-full
            [&::-webkit-slider-thumb]:shadow-lg
            [&::-webkit-slider-thumb]:cursor-pointer
            [&::-webkit-slider-thumb]:transition-transform
            [&::-webkit-slider-thumb]:active:scale-110
            [&::-webkit-slider-thumb]:border-2
            [&::-webkit-slider-thumb]:border-white
            [&::-webkit-slider-thumb]:-mt-[6px]
            [&::-moz-range-thumb]:relative
            [&::-moz-range-thumb]:z-20
            [&::-moz-range-thumb]:w-5
            [&::-moz-range-thumb]:h-5
            [&::-moz-range-thumb]:bg-[#8b5cf6]
            [&::-moz-range-thumb]:rounded-full
            [&::-moz-range-thumb]:border-2
            [&::-moz-range-thumb]:border-white
            [&::-moz-range-thumb]:cursor-pointer
            [&:focus]:outline-none
          "
        />
      </div>

      <div className="flex justify-between mt-1">
        <span className="text-xs text-[#71717a]">{min}{suffix}</span>
        <span className="text-xs text-[#71717a]">{max}{suffix}</span>
      </div>
    </div>
  );
}