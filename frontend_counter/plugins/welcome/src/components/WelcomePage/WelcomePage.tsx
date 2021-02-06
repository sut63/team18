import React, { FC } from 'react';
import { Typography, Grid, Link } from '@material-ui/core';
import Divider from '@material-ui/core/Divider';
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
import boss from '../../image/FullArtKyaru.png'
import best from '../../image/FullArt6Kyaru.png'
import tong from '../../image/FullArtSummerKyaru.png'
import film from '../../image/FullArtNewYearKyaru.png'
import ta from '../../image/InitialKyaru2.png'

const HeaderCustom = {
  minHeight: '50px',
};

const useStyles = makeStyles(theme => ({
  root: {
    maxWidth: 345,
  },
  divider: {
    margin: theme.spacing(2, 0),
  },
}));

export type ProfileProps = {
  name: string;
  id: string;
  system: string;
  img: string;
  link: string;
};

export function CardTeam({ name, id, system, img, link }: ProfileProps) {
  const classes = useStyles();
  return (
    <Grid item xs={12} md={3}>
      <Card className={classes.root}>
        <CardActionArea>
          <Link
            href={link}
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
  const classes = useStyles();
  return (
    <Page theme={pageTheme.website}>
      <Header style={HeaderCustom} title={`ระบบจองห้องพัก`}></Header>
      <Content>
        <ContentHeader title="รายการระบบย่อย Sprint #1"></ContentHeader>
        <Grid container>
          <CardTeam name={"นายภควัฒน์ เพ็ญภูมิ"} id={"B6111465"} system={"ระบบย่อย Check In"} img="https://www.akerufeed.com/wp-content/uploads/2020/08/%E0%B9%81%E0%B8%A1%E0%B8%A7%E0%B8%AA%E0%B9%89%E0%B8%A1-1.jpg?fbclid=IwAR10_OLyOO2tnwJZl0ye_Hz9nxLZCWTUPekTKjh9MU_57p8SJekUXkShZ-g" link="/CheckIn"></CardTeam>
          <CardTeam name={"นายจิรัฎฐ์ เมธากุลกิติวัฒน์"} id={"B6112479"} system={"ระบบย่อย Furniture Detail"} img={best} link="/FurnitureDetail"></CardTeam>
          <CardTeam name={"นายชัยฤทธิ์ แซ่ฮึง"} id={"B6131081"} system={"ระบบย่อยข้อมูลห้องพัก"} img={ta} link="/dataroom"></CardTeam>
          <CardTeam name={"นายอิสรภาพ ศรีบุตรตา"} id={"B6131685"} system={"ระบบย่อย Check Out"} img={film} link="/checkout"></CardTeam>
        </Grid>
        <Grid item xs={12}>
          <Divider className={classes.divider} />
          <Typography variant="h4" gutterBottom>
              รายการระบบย่อย Sprint #2
          </Typography>
        </Grid>
        <Grid container>
          <CardTeam name={"นายสิรภพ นิธิภัทรชัย"} id={"B6108380"} system={"ระบบย่อยค้นหาใบแจ้งซ่อม"} img={boss} link="/SearchFixRoom"></CardTeam>
          <CardTeam name={"นายภควัฒน์ เพ็ญภูมิ"} id={"B6111465"} system={"ระบบย่อยค้นหา Check In"} img="https://www.akerufeed.com/wp-content/uploads/2020/08/%E0%B9%81%E0%B8%A1%E0%B8%A7%E0%B8%AA%E0%B9%89%E0%B8%A1-1.jpg?fbclid=IwAR10_OLyOO2tnwJZl0ye_Hz9nxLZCWTUPekTKjh9MU_57p8SJekUXkShZ-g{karyl3}" link="/SearchCheckIn"></CardTeam>
          <CardTeam name={"นายจิรัฎฐ์ เมธากุลกิติวัฒน์"} id={"B6112479"} system={"ระบบย่อยค้นหาเฟอร์นิเจอร์"} img={best} link="/SearchFurniture"></CardTeam>
          <CardTeam name={"นายวิษณุ สันติภาษิต"} id={"B6115234"} system={"ระบบย่อยค้นหาใบจอง"} img={tong} link="/SearchReserveRoom"></CardTeam>
          <CardTeam name={"นายอิสรภาพ ศรีบุตรตา"} id={"B6131685"} system={"ระบบย่อยค้นหา Check Out"} img={film} link="/SearchCheckout"></CardTeam>
        </Grid>
      </Content>
    </Page>
  );
};

export default WelcomePage;
