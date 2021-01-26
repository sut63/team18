import { createPlugin } from '@backstage/core';
import FixRoom from './components/FixRoom';
import Reserve from './components/ReserveRoom'
import RoomList from './components/RoomList'
import SignIn from './components/SignIn'
import SearchDataRoom from './components/SearchDataRoom'
import { Cookies } from './Cookie'

var ck = new Cookies()
var role = ck.GetRole()

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    if(role != "customer"){
      router.registerRoute('/', SignIn);
      router.registerRoute('/reserve', SignIn);
      router.registerRoute('/fix', SignIn);
      router.registerRoute('/sign_in', SignIn);
      router.registerRoute('/searchdataroom', SignIn);
    }else{
      router.registerRoute('/', RoomList);
      router.registerRoute('/reserve', Reserve);
      router.registerRoute('/fix', FixRoom);
      router.registerRoute('/sign_in', SignIn);
      router.registerRoute('/searchdataroom', SearchDataRoom);
    }
  },
});

