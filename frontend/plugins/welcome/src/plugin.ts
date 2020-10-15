import { createPlugin } from '@backstage/core';
import Books from './components/Books';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', Books);
  },
});
