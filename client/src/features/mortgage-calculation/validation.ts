export interface ValidationErrors {
    propertyPrice?: string;
    downPaymentAmount?: string;
    matCapitalAmount?: string;
    interestRate?: string;
}

export const validateForm = (formData: {
    propertyPrice: number;
    downPaymentAmount: number;
    matCapitalAmount: number;
    interestRate: number;
}) => {
    const errors: ValidationErrors = {};

    if (formData.downPaymentAmount > formData.propertyPrice) {
        errors.downPaymentAmount = 'Первоначальный взнос не может быть больше стоимости недвижимости';
    }

    if (formData.downPaymentAmount < 0) {
        errors.downPaymentAmount = 'Первоначальный взнос не может быть отрицательным';
    }

    if (formData.matCapitalAmount > formData.propertyPrice - formData.downPaymentAmount) {
        errors.matCapitalAmount = 'Материнский капитал не может быть больше суммы кредита';
    }

    if (formData.interestRate < 1 || formData.interestRate > 30) {
        errors.interestRate = 'Процентная ставка должна быть от 1% до 30%';
    }

    if (formData.propertyPrice < 100000) {
        errors.propertyPrice = 'Стоимость недвижимости должна быть не менее 100 000 ₽';
    }

    return errors;
};