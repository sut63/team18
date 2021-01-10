import { createPlugin } from '@backstage/core';
import FixRoom from './components/FixRoom';
import Reserve from './components/ReserveRoom'
import RoomList from './components/RoomList'
import SignIn from './components/SignIn'
import { Cookies } from './Cookie'

var ck = new Cookies()
var cookie = ck.GetCookie()

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    if(cookie == undefined){
      router.registerRoute('/', SignIn);
      router.registerRoute('/reserve', SignIn);
      router.registerRoute('/fix', SignIn);
      router.registerRoute('/sign_in', SignIn);
    }else{
      router.registerRoute('/', RoomList);
      router.registerRoute('/reserve', Reserve);
      router.registerRoute('/fix', FixRoom);
      router.registerRoute('/sign_in', SignIn);
    }
  },
});

