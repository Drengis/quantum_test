import { Suspense } from 'react';
import SchedulePage from '@/pages/schedule/ui/SchedulePage';

export default function Schedule() {
  return (
    <Suspense fallback={
      <div className="flex items-center justify-center min-h-[50vh]">
        <div className="w-12 h-12 rounded-full bg-gradient-to-r from-[#8b5cf6] to-[#7c3aed] animate-pulse-subtle" />
      </div>
    }>
      <SchedulePage />
    </Suspense>
  );
}