import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import SignIn from './components/SignIn'
import Dataroom from './components/DataRoom'
import CheckIn from './components/CheckIn'
import checkout from './components/Checkout'
import FurnitureDetail from './components/FurnitureDetail'
import SearchReserveRoom from './components/SearchReserveRoom'
import SearchCheckIn from './components/SearchCheckIn'
import { Cookies } from './Cookie'
import SearchCheckout from './components/SearchCheckout'
import SearchFixRoom from './components/SearchFixRoom'
import searchfurniture from './components/SearchFurniture'
var ck = new Cookies()
var role = ck.GetRole()

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    if(role != "counterStaff"){
      router.registerRoute('/', SignIn);
      router.registerRoute('/WelcomePage', SignIn);
      router.registerRoute('/dataroom', SignIn);
      router.registerRoute('/CheckIn', SignIn);
      router.registerRoute('/checkout', SignIn);
      router.registerRoute('/FurnitureDetail', SignIn);
      router.registerRoute('/SearchCheckIn', SignIn);
      router.registerRoute('/SearchCheckout', SignIn);    
    }else{
      router.registerRoute('/', WelcomePage);
      router.registerRoute('/WelcomePage', WelcomePage);
      router.registerRoute('/dataroom', Dataroom);
      router.registerRoute('/CheckIn', CheckIn);
      router.registerRoute('/checkout', checkout);
      router.registerRoute('/FurnitureDetail', FurnitureDetail);
      router.registerRoute('/SearchReserveRoom', SearchReserveRoom);
      router.registerRoute('/SearchCheckIn', SearchCheckIn);
      router.registerRoute('/SearchCheckout', SearchCheckout);
      router.registerRoute('/SearchFixRoom', SearchFixRoom);
      router.registerRoute('/SearchFurniture', searchfurniture);    
    }
  },
});
