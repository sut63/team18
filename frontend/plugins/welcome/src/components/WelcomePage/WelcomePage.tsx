import React, { FC } from 'react';
import { Typography, Grid, Link } from '@material-ui/core';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import karyl6 from '../../image/FullArt6Kyaru.png'
import karylsummer from '../../image/FullArtSummerKyaru.png'
import karylnewyear from '../../image/FullArtNewYearKyaru.png'

const HeaderCustom = {
  minHeight: '50px',
};

const useStyles = makeStyles({
  root: {
    maxWidth: 345,
  },
});

export type ProfileProps = {
  name: string; 
  id: string;
  system: string;
  img: string;
  link: string;
};

export function CardTeam({ name, id, system, img, link}: ProfileProps) {
  const classes = useStyles();
  return (
    <Grid item xs={12} md={3}>
      <Card className={classes.root}>
        <CardActionArea>
        <Link
        href = {link}
        >
          <CardMedia
            component="img"
            height="140"
            image={img}  
          />
           </Link>
            <Typography gutterBottom variant="h5" component="h4">
              {system}
            </Typography>
            <Typography gutterBottom variant="h6" component="h2">
              {id} {name}
            </Typography>
        </CardActionArea>
      </Card>
    </Grid>
  );
}

const WelcomePage: FC<{}> = () => {
  const profile = { givenName: 'ระบบจองห้องพัก' };
  return (
    <Page theme={pageTheme.website}>
      <Header style={HeaderCustom} title={`ยินดีต้อนรับสู่ ${profile.givenName || 'to Backstage'}`}
       subtitle="สำหรับลูกค้า"></Header>
      <Content>
        <ContentHeader title="รายการระบบย่อย"></ContentHeader>
        <Grid container>
          <CardTeam name={"นายวิษณุ สันติภาษิต"} id={"B6115234"} system={"ระบบย่อยจองห้องพัก"} img = {karyl6} link = "/reserve"></CardTeam>
          <CardTeam name={"นายสิรภพ นิธิภัทรชัย"} id={"B6108380"} system={"ระบบย่องแจ้งซ่อม"} img = {karylsummer} link = "/fix"></CardTeam>
          <CardTeam name={"นายชัยฤทธิ์ แซ่ฮึง"} id={"B6131081"} system={"ระบบย่อยค้นหาห้องพัก"} img = {karylnewyear} link = "/searchdataroom"></CardTeam>
        </Grid>
      </Content>
    </Page>
  );
};

export default WelcomePage;
