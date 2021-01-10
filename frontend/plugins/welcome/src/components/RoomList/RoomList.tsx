import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import TablesRoom from '../TableRoom';
import Button from '@material-ui/core/Button';
 
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';
 
const RoomList: FC<{}> = () => {
 const profile = { givenName: 'ระบบจองห้องพัก' };
 
 return (
   <Page theme={pageTheme.home}>
     <Header
       title={`ยินดีต้อนรับสู่ ${profile.givenName || 'to Backstage'}`}
       subtitle="Some quick intro and links."
     ></Header>
     <Content>
       <ContentHeader title="ไปรายการห้องพัก">
         <Link component={RouterLink} to="/reserve">
           <Button variant="contained" color="primary">
             จองห้อง
           </Button>
         </Link>
       </ContentHeader>
       <TablesRoom></TablesRoom>
     </Content>
   </Page>
 );
};
 
export default RoomList;