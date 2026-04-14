'use client';

import { useEffect, ReactNode } from 'react';
import { init, viewport, miniApp } from '@tma.js/sdk';

export function TelegramProvider({ children }: { children: ReactNode }) {
  useEffect(() => {
    const run = async () => {
      try {

        await init();

        if (viewport.mount.isAvailable()) {
          await viewport.mount();
          viewport.expand?.();
        }

        if (miniApp.mount.isAvailable()) {
          await miniApp.mount();
          miniApp.ready();
        }

      } catch (e) {
        console.error('TMA Init Error:', e);
      }
    };

    run();
  }, []);

  return <>{children}</>;
}