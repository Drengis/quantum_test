'use client';

import { useUnit } from 'effector-react';
import { useRouter } from 'next/navigation';
import { useInitUser } from '@/entities/user/model/useInitUser';
import { Slider } from '@/shared/ui/Slider';
import { Input } from '@/shared/ui/Input';
import { Select } from '@/shared/ui/Select';
import {
  $formData,
  $isLoading,
  $isPolling,
  $result,
  $error,
  $validationErrors,
  updateFormData,
  submitForm,
  validateFormEvent,
} from '@/features/mortgage-calculation/model';

const propertyTypes = [
  { value: 'apartment_in_new_building', label: 'Квартира в новостройке' },
  { value: 'apartment_in_secondary_building', label: 'Квартира на вторичном' },
  { value: 'house', label: 'Жилой дом' },
  { value: 'house_with_land_plot', label: 'Дом с участком' },
  { value: 'land_plot', label: 'Земельный участок' },
];

const formatPrice = (value: number) => {
  return new Intl.NumberFormat('ru-RU').format(value);
};

export default function HomePage() {
  const router = useRouter();
  const { user, isLoading: userLoading } = useInitUser();

  const formData = useUnit($formData);
  const isLoading = useUnit($isLoading);
  const isPolling = useUnit($isPolling);
  const result = useUnit($result);
  const error = useUnit($error);
  const validationErrors = useUnit($validationErrors);

  const updateForm = useUnit(updateFormData);
  const onSubmit = useUnit(submitForm);
  const validate = useUnit(validateFormEvent);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (!user?.id) return;

    validate();
    onSubmit({ userId: user.id });
  };

  if (userLoading) {
    return (
      <div className="min-h-[50vh] flex items-center justify-center">
        <div className="w-12 h-12 rounded-full bg-gradient-to-r from-[#8b5cf6] to-[#7c3aed] animate-pulse-subtle" />
      </div>
    );
  }

  if (!user) {
    return (
      <div className="flex items-center justify-center min-h-[50vh]">
        <div className="text-red-400 animate-fade-in">Ошибка инициализации</div>
      </div>
    );
  }

  return (
    <div className="p-4 pb-20 max-w-md mx-auto">
      <div className="animate-fade-in-up delay-0">
        <h1 className="text-2xl font-bold mb-2 bg-gradient-to-r from-white to-[#a1a1aa] bg-clip-text text-transparent">
          Ипотека
        </h1>
      </div>

      <div className="animate-fade-in-up delay-100">
        <p className="text-[#a1a1aa] mb-6">
          Привет, {user.first_name}!
        </p>
      </div>

      <form onSubmit={handleSubmit} className="space-y-5 stagger-children">
        <div>
          <Select
            label="Тип недвижимости"
            value={formData.propertyType}
            onChange={(value) => updateForm({ propertyType: value })}
            options={propertyTypes}
          />
        </div>

        <div>
          <Slider
            label="Стоимость недвижимости"
            value={formData.propertyPrice}
            onChange={(v) => updateForm({ propertyPrice: v })}
            min={500000}
            max={50000000}
            step={100000}
            suffix="₽"
            formatValue={formatPrice}
          />
          {validationErrors.propertyPrice && (
            <p className="text-red-400 text-xs mt-1 animate-shake">
              {validationErrors.propertyPrice}
            </p>
          )}
        </div>

        <div>
          <Slider
            label="Первоначальный взнос"
            value={formData.downPaymentAmount}
            onChange={(v) => updateForm({ downPaymentAmount: v })}
            min={0}
            max={formData.propertyPrice}
            step={50000}
            suffix="₽"
            formatValue={formatPrice}
          />
          {validationErrors.downPaymentAmount && (
            <p className="text-red-400 text-xs mt-1 animate-shake">
              {validationErrors.downPaymentAmount}
            </p>
          )}
        </div>

        <div>
          <Slider
            label="Срок ипотеки"
            value={formData.mortgageTermYears}
            onChange={(v) => updateForm({ mortgageTermYears: v })}
            min={1}
            max={30}
            step={1}
            suffix="лет"
          />
        </div>

        <div>
          <Slider
            label="Процентная ставка"
            value={formData.interestRate}
            onChange={(v) => updateForm({ interestRate: v })}
            min={1}
            max={20}
            step={0.1}
            suffix="%"
          />
          {validationErrors.interestRate && (
            <p className="text-red-400 text-xs mt-1 animate-shake">
              {validationErrors.interestRate}
            </p>
          )}
        </div>

        <label className="flex items-center gap-3 p-3 bg-[#1c1c1f] rounded-xl border border-[#3f3f46] cursor-pointer transition-all hover:border-[#8b5cf6] hover-glow">
          <input
            type="checkbox"
            checked={formData.matCapitalIncluded}
            onChange={(e) => updateForm({ matCapitalIncluded: e.target.checked })}
            className="w-5 h-5 accent-[#8b5cf6]"
          />
          <span className="text-sm">Материнский капитал</span>
        </label>

        {formData.matCapitalIncluded && (
          <div className="animate-slide-down">
            <Input
              label="Сумма мат. капитала"
              value={formData.matCapitalAmount === 0 ? '' : String(formData.matCapitalAmount)}
              onChange={(v) => updateForm({ matCapitalAmount: Number(v) || 0 })}
              type="number"
              suffix="₽"
            />
            {validationErrors.matCapitalAmount && (
              <p className="text-red-400 text-xs mt-1 animate-shake">
                {validationErrors.matCapitalAmount}
              </p>
            )}
          </div>
        )}

        <button
          type="submit"
          disabled={isLoading || isPolling}
          className="w-full bg-[#8b5cf6] hover:bg-[#7c3aed] active:scale-[0.98] disabled:opacity-50 text-white font-semibold py-4 px-6 rounded-xl transition-all hover-lift"
        >
          {isLoading ? 'Отправка...' : isPolling ? 'Расчёт...' : 'Рассчитать'}
        </button>

        {error && (
          <div className="p-3 bg-red-500/10 border border-red-500/30 rounded-lg text-red-400 text-sm animate-shake">
            {error}
          </div>
        )}
      </form>

      {result && result.status === 'pending' && (
        <div className="mt-6 p-5 bg-[#1c1c1f] border border-[#3f3f46] rounded-2xl text-center animate-fade-in-scale">
          <div className="w-8 h-8 border-2 border-[#8b5cf6] border-t-transparent rounded-full animate-spin mx-auto mb-3" />
          <p className="text-[#a1a1aa]">Идёт расчёт ипотеки...</p>
          <p className="text-[#71717a] text-sm mt-2 animate-pulse-subtle">Пожалуйста, подождите</p>
        </div>
      )}

      {result && result.status === 'completed' && (
        <div className="mt-6 p-5 bg-[#1c1c1f] border border-[#3f3f46] rounded-2xl result-enter">
          <h2 className="text-lg font-semibold mb-4">Результат</h2>
          <div className="space-y-3 stagger-children">
            <div className="flex justify-between py-2 border-b border-[#3f3f46]">
              <span className="text-[#a1a1aa]">Ежемесячный платёж</span>
              <span className="font-semibold text-[#8b5cf6]">{result.monthlyPayment} ₽</span>
            </div>
            <div className="flex justify-between py-2 border-b border-[#3f3f46]">
              <span className="text-[#a1a1aa]">Общая сумма</span>
              <span className="font-semibold">{result.totalPayment} ₽</span>
            </div>
            <div className="flex justify-between py-2 border-b border-[#3f3f46]">
              <span className="text-[#a1a1aa]">Переплата</span>
              <span className="font-semibold text-red-400">{result.totalOverpaymentAmount} ₽</span>
            </div>
            <div className="flex justify-between py-2 border-b border-[#3f3f46]">
              <span className="text-[#a1a1aa]">Налоговый вычет</span>
              <span className="font-semibold text-green-400">{result.possibleTaxDeduction} ₽</span>
            </div>
            <div className="flex justify-between py-2">
              <span className="text-[#a1a1aa]">Мин. доход</span>
              <span className="font-semibold">{result.recommendedIncome} ₽</span>
            </div>
          </div>
          <button
            onClick={() => router.push(`/schedule?id=${result.id}`)}
            className="mt-4 w-full bg-[#252529] hover:bg-[#3f3f46] active:scale-[0.98] text-white font-medium py-3 px-6 rounded-xl transition-all hover-lift"
          >
            График платежей →
          </button>
        </div>
      )}
    </div>
  );
}