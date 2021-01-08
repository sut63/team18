import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import SignIn from './components/SignIn'
import Dataroom from './components/DataRoom'
import CheckIn from './components/CheckIn'


export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/signin', SignIn);
    router.registerRoute('/dataroom', Dataroom);
    router.registerRoute('/CheckIn', CheckIn);
  },
});
