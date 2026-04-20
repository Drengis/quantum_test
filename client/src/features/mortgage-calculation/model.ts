import { createStore, createEvent, createEffect, sample } from 'effector';
import { api } from '@/shared/api';
import { validateForm, ValidationErrors } from './validation';

import { MortgageProfile as FormData, MortgageResult } from '@/entities/mortgage/model/types';

const initialState: FormData = {
    propertyType: 'apartment_in_new_building',
    propertyPrice: 5000000,
    downPaymentAmount: 1000000,
    mortgageTermYears: 20,
    interestRate: 8.5,
    matCapitalIncluded: false,
    matCapitalAmount: 0,
};

export const updateFormData = createEvent<Partial<FormData>>();
export const resetForm = createEvent();
export const submitForm = createEvent<{ userId: string }>();
export const validateFormEvent = createEvent();

export const submitMortgageCalculationFx = createEffect<
    { userId: string; formData: FormData },
    { id: string }
>(async ({ userId, formData }) => {
    const response = await api.post<{ id: string; status: string }>('/mortgage-profiles', {
        user_id: userId,
        propertyType: formData.propertyType,
        propertyPrice: formData.propertyPrice,
        downPaymentAmount: formData.downPaymentAmount,
        mortgageTermYears: formData.mortgageTermYears,
        interestRate: formData.interestRate,
        matCapitalIncluded: formData.matCapitalIncluded,
        matCapitalAmount: formData.matCapitalIncluded ? formData.matCapitalAmount : null,
    });
    return { id: response.id };
});

export const getMortgageResultFx = createEffect<
    { id: string },
    MortgageResult
>(async ({ id }) => {
    const result = await api.get<MortgageResult>(`/mortgage-profiles/${id}`);
    return result;
});

export const $formData = createStore<FormData>(initialState)
    .on(updateFormData, (state, payload) => ({ ...state, ...payload }))
    .reset(resetForm);

export const $validationErrors = createStore<ValidationErrors>({});
export const $isLoading = submitMortgageCalculationFx.pending;
export const $isPolling = createStore<boolean>(false);
export const $result = createStore<MortgageResult | null>(null);
export const $error = createStore<string | null>(null);

sample({
    clock: updateFormData,
    source: $formData,
    fn: (formData) => validateForm(formData),
    target: $validationErrors,
});

sample({
    clock: validateFormEvent,
    source: $formData,
    fn: (formData) => validateForm(formData),
    target: $validationErrors,
});

const pollCalculationFx = createEffect<{ id: string }, MortgageResult>(async ({ id }) => {
    const maxAttempts = 30;
    const interval = 2000;

    for (let i = 0; i < maxAttempts; i++) {
        await new Promise(resolve => setTimeout(resolve, interval));
        const data = await api.get<MortgageResult>(`/mortgage-profiles/${id}`);

        if (data.status === 'completed') {
            return data;
        }
    }
    throw new Error('Расчёт занимает больше времени, попробуйте позже');
});

sample({
    clock: submitForm,
    source: { formData: $formData, errors: $validationErrors },
    filter: ({ errors }) => Object.keys(errors).length === 0,
    fn: ({ formData }, clockData) => ({
        userId: clockData.userId,
        formData
    }),
    target: submitMortgageCalculationFx,
});

sample({
    clock: submitForm,
    source: $validationErrors,
    filter: (errors) => Object.keys(errors).length > 0,
    fn: () => 'Пожалуйста, исправьте ошибки в форме',
    target: $error,
});

sample({
    clock: submitMortgageCalculationFx.doneData,
    fn: () => true,
    target: $isPolling,
});

sample({
    clock: submitMortgageCalculationFx.doneData,
    fn: ({ id }) => ({ id }),
    target: pollCalculationFx,
});

sample({
    clock: pollCalculationFx.doneData,
    target: $result,
});

sample({
    clock: pollCalculationFx.doneData,
    fn: () => false,
    target: $isPolling,
});

sample({
    clock: pollCalculationFx.failData,
    fn: (error) => error.message,
    target: $error,
});

sample({
    clock: pollCalculationFx.failData,
    fn: () => false,
    target: $isPolling,
});

sample({
    clock: submitForm,
    fn: () => null,
    target: $error,
});

sample({
    clock: submitForm,
    fn: () => null,
    target: $result,
});

sample({
    clock: submitMortgageCalculationFx.doneData,
    fn: () => ({}),
    target: $validationErrors,
});