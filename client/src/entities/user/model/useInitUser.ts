'use client';

import { useEffect } from 'react';
import { useStore } from 'effector-react';
import { initData } from '@tma.js/sdk';
import { $user, $isUserLoading, initUserFx } from './store';
import type { TelegramUser } from './types';

const MOCK_USER: TelegramUser = {
  tg_id: 123456789,
  username: 'testuser',
  first_name: 'Test',
  last_name: 'User',
  language_code: 'ru',
};

const isDev = process.env.NODE_ENV === 'development';

export function useInitUser() {
  const user = useStore($user);
  const isLoading = useStore($isUserLoading);

  useEffect(() => {
    if (user) return;

    try {
      const initDataUser = initData.user();

      if (initDataUser?.id && initDataUser?.first_name) {
        initUserFx({
          tg_id: initDataUser.id,
          username: initDataUser.username,
          first_name: initDataUser.first_name,
          last_name: initDataUser.last_name,
          language_code: initDataUser.language_code,
        });
        return;
      }
    } catch (e) {
      console.warn('Not in Telegram or initData not available');
    }

    if (isDev) {
      initUserFx(MOCK_USER);
    }
  }, [user]);

  return { user, isLoading };
}