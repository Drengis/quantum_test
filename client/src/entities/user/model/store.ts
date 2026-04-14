import { createEffect, createStore } from 'effector';
import { api } from '@/shared/api';
import type { User, CreateUserRequest, TelegramUser } from './types';

export const $user = createStore<User | null>(null);
export const $isUserLoading = createStore<boolean>(true);

export const initUserFx = createEffect(async (telegramUser: TelegramUser) => {
  const payload: CreateUserRequest = {
    tg_id: String(telegramUser.tg_id),
    username: telegramUser.username,
    first_name: telegramUser.first_name || 'User',
    last_name: telegramUser.last_name,
    lang_code: telegramUser.language_code,
  };

  const response = await api.post<User>('/user', payload);
  return response;
});

$user.on(initUserFx.doneData, (_, user) => user);
$isUserLoading.on(initUserFx.finally, () => false);