export interface MortgageProfile {
  propertyType: string;
  propertyPrice: number;
  downPaymentAmount: number;
  matCapitalAmount: number;
  matCapitalIncluded: boolean;
  mortgageTermYears: number;
  interestRate: number;
}

export interface MortgageResult {
  id: string;
  status: string;
  monthlyPayment: string;
  totalPayment: string;
  totalOverpaymentAmount: string;
  possibleTaxDeduction: string;
  savingsDueMotherCapital: string;
  recommendedIncome: string;
  mortgagePaymentSchedule?: Record<string, Record<string, {
    totalPayment: string;
    repaymentOfMortgageBody: string;
    repaymentOfMortgageInterest: string;
    mortgageBalance: string;
  }>>;
}

export interface PropertyType {
  value: string;
  label: string;
}

export const propertyTypes: PropertyType[] = [
  { value: 'apartment_in_new_building', label: 'Квартира в новостройке' },
  { value: 'apartment_in_secondary_building', label: 'Квартира на вторичном рынке' },
  { value: 'house', label: 'Жилой дом' },
  { value: 'house_with_land_plot', label: 'Дом с земельным участком' },
  { value: 'land_plot', label: 'Земельный участок' },
];
