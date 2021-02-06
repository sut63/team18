import React, { FC, useEffect } from 'react';
import { Theme, makeStyles, withStyles, createStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2'; // alert
import {
  MuiPickersUtilsProvider,
  KeyboardTimePicker,
  KeyboardDatePicker,
} from '@material-ui/pickers';
import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Avatar,
  Button,
  Badge,
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntDataRoom } from '../../api/models/EntDataRoom';
import { EntFurniture } from '../../api/models/EntFurniture';
import { EntCounterStaff } from '../../api/models/EntCounterStaff';
import { Cookies } from '../../Cookie'
import Snackbar from '@material-ui/core/Snackbar';
import Alert from '@material-ui/lab/Alert';
import accountImg from '../../image/account.jpg'

// header css
const HeaderCustom = {
  minHeight: '50px',
};

const StyledBadge = withStyles((theme: Theme) =>
  createStyles({
    badge: {
      backgroundColor: '#44b700',
      color: '#44b700',
      boxShadow: `0 0 0 2px ${theme.palette.background.paper}`,
      '&::after': {
        position: 'absolute',
        top: 0,
        left: 0,
        width: '100%',
        height: '100%',
        borderRadius: '50%',
        animation: '$ripple 1.2s infinite ease-in-out',
        border: '1px solid currentColor',
        content: '""',
      },
    },
    '@keyframes ripple': {
      '0%': {
        transform: 'scale(.8)',
        opacity: 1,
      },
      '100%': {
        transform: 'scale(2.4)',
        opacity: 0,
      },
    },
  }),
)(Badge);

// css style
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    width: 300,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: 300,
  },
}));

interface FurnitureDetail {
  Dataroom: number;
  Furniture: number;
  CounterStaff: number;
  DateAdd: Date;
  FurnitureAmount: number;
  FurnitureColour: string;
  Detail: string;
  // create_by: number;
}

interface room {
  StatusRoom: number;
  // create_by: number;
}

const FurnitureDetail: FC<{}> = () => {
  const classes = useStyles();
  const api = new DefaultApi();

  // ดึงคุกกี้
  var ck = new Cookies()
  var cookieName = ck.GetCookie()
  var cookieID = ck.GetID()

  //อันหลักสำหรับสร้าง การ furniture_detail
  const [furniture_detail, setFurnitureDetail] = React.useState<Partial<FurnitureDetail>>({});
  furniture_detail.FurnitureAmount = Number(furniture_detail.FurnitureAmount)
  // data room List
  const [dataroom, setdataroom] = React.useState<EntDataRoom[]>([]);
  const getDataRoom = async () => {
    const res = await api.listDataroom({ limit: 10, offset: 0 });
    setdataroom(res);
  };


  // furniture List
  const [furnitures, setFurniture] = React.useState<EntFurniture[]>([]);
  const getFurniture = async () => {
    const res = await api.listFurniture({ limit: 10, offset: 0 });
    setFurniture(res);
  };

  const [FurnitureAmountError, setFurnitureAmountError] = React.useState('');
  const [FurnitureColourError, setFurnitureColourError] = React.useState('');
  const [DetailError, setDetailError] = React.useState('');

  //counterstaff
  const [counterstaffs, setCounterStaffs] = React.useState<EntCounterStaff>();
  const getCounterStaffs = async () => {
    const res = await api.getCounterStaff({ id: Number(cookieID) });
    setCounterStaffs(res);
  };


  //set time
  const [dateAdd, setDateAdd] = React.useState<any>(0)

  const validateFurnitureColour = (val: string) => {
    return val.length > 10 ? false : true;
  }

  const validateFurnitureAmount = (val: Number) => {
    return val > 0 && val <= 10 ? true : false;
  }

  // ฟังก์ชั่นสำหรับ validate รหัสนักศึกษา
  const validateDetail = (val: string) => {
    return val.length > 50 ? false : true;
  }
  //checkpatten
  const checkPattern = (id: string, value: string) => {
    switch (id) {
      case 'FurnitureColour':
        validateFurnitureColour(value) ? setFurnitureColourError('') : setFurnitureColourError('ห้ามเกิน 10 ตัวอักษร');
        return;
      case 'FurnitureAmount':
        validateFurnitureAmount(Number(value)) ? setFurnitureAmountError('') : setFurnitureAmountError('เฟอร์นิเจอร์ไม่เกิน 10 ชิ้น');
        return;
      case 'Detail':
        validateDetail(value) ? setDetailError('') : setDetailError('ห้ามเกิน 50 ตัวอักษร')
        return;
      default:
        return;
    }
  }

  const checkerror = (s: string) => {
    switch (s) {
      case 'furniture_colour':
        setError("รายละเอียดมีความยาวมากเกินไป")
        return;
      case 'furniture_amount':
        setError("จำนวนเฟอร์นิเจอร์ไม่ถูกต้อง")
        return;
      case 'detail':
        setError("รายละเอียดมีความยาวมากเกินไป")
        return;
      default:
        return;
    }
  };

  // alert setting
  const [open, setOpen] = React.useState(false);
  const [fail, setFail] = React.useState(false);
  const [errors, setError] = React.useState(String);
  //close alert 
  const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
    if (reason === 'clickaway') {
      return;
    }
    setFail(false);
    setOpen(false);
  };


  // Lifecycle Hooks
  useEffect(() => {
    getCounterStaffs();
    getFurniture();
    getDataRoom();
  }, []);


  useEffect(() => {
    setFurnitureDetail({ ...furniture_detail, ['CounterStaff']: counterstaffs?.id });
  }, [counterstaffs]);

  // set data to object furniture_detail
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: any }>,
  ) => {
    const name = event.target.name as keyof typeof FurnitureDetail;
    const { value } = event.target;
    setFurnitureDetail({ ...furniture_detail, [name]: value });
    const validateValue = value.toString()
    checkPattern(name, validateValue)
    console.log(furniture_detail);
  };

  // set data for time in
  const handleChangeDateAdd = (
    event: React.ChangeEvent<{ name?: string; value: any }>,
  ) => {
    const name = event.target.name as keyof typeof FurnitureDetail;
    const { value } = event.target;
    const time = value + ":00+07:00"
    setFurnitureDetail({ ...furniture_detail, [name]: time });
    setDateAdd(value);
    console.log(furniture_detail);
  };


  // clear input form
  function clear() {
    setFurnitureDetail({});
    setDateAdd({});
  }


  // clear cookies
  function Clears() {
    ck.ClearCookie()
    window.location.reload(false)
  }

  // function save data
  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/furnituredetails';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(furniture_detail),
    };

    console.log(furniture_detail); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data.status);
        if (data.status === true) {
          clear();
          setOpen(true);
        } else {
          checkerror(data.error.Name)
          setFail(true);
        }
      });
  }



  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`ระบบจัดการเฟอร์นิเจอร์`}>
        <StyledBadge
          overlap="circle"
          anchorOrigin={{
            vertical: 'bottom',
            horizontal: 'right',
          }}
          variant="dot"
        >
          <Avatar alt="Remy Sharp" src={accountImg} />
        </StyledBadge>
        <div style={{ marginLeft: 10, marginRight: 20 }}>{cookieName}</div>
        <Button
          variant="outlined"
          color="secondary"
          size="large"
          onClick={Clears}
        >
          Logout
        </Button>
      </Header>
      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>พนักงาน</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  disabled
                  name="CounterStaff"
                  variant="outlined"
                  value={counterstaffs?.name}
                />
              </FormControl>
            </Grid>



            <Grid item xs={3}>
              <div className={classes.paper}>room</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกห้องพัก</InputLabel>
                <Select
                  name="Dataroom"
                  value={furniture_detail.Dataroom || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {dataroom.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.roomnumber}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>furniture</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกเฟอร์นิเจอร์</InputLabel>
                <Select
                  name="Furniture"
                  value={furniture_detail.Furniture || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {furnitures.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.furnitureName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>



            <Grid item xs={3}>
              <div className={classes.paper}>Detail</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  error={DetailError ? true : false}
                  helperText={DetailError}
                  name="Detail"
                  label="รายละเอียด"
                  variant="outlined"
                  value={furniture_detail.Detail || ''}
                  onChange={handleChange}
                />
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>Furniture Colour</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  error={FurnitureColourError ? true : false}
                  helperText={FurnitureColourError}
                  name="FurnitureColour"
                  label="สีของเฟอร์นิเจอร์"
                  variant="outlined"
                  value={furniture_detail.FurnitureColour || ''}
                  onChange={handleChange}
                />
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>จำนวนเฟอร์นิเจอร์</div>
            </Grid>
            <Grid item xs={9}>
              <TextField
                error={FurnitureAmountError ? true : false}
                helperText={FurnitureAmountError}
                type={"number"}
                value={furniture_detail.FurnitureAmount || ''}
                label="จำนวนเฟอร์นิเจอร์"
                name="FurnitureAmount"
                variant="outlined"
                onChange={handleChange} />
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>เลือกวันที่เพิ่มเฟอร์นิเจอร์</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  name="Dateadd"
                  type="datetime-local"
                  value={dateAdd}
                  defaultValue="2020-12-31"
                  onChange={handleChangeDateAdd}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />
              </FormControl>
            </Grid>


            <Grid item xs={3}></Grid>
            <Grid item xs={9}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={save}
              >
                บันทึก
              </Button>
            </Grid>

          </Grid>
          <Snackbar open={open} autoHideDuration={6000} onClose={handleClose}>
            <Alert onClose={handleClose} severity="success">
              บันทึกสำเร็จ!
        </Alert>
          </Snackbar>

          <Snackbar open={fail} autoHideDuration={6000} onClose={handleClose}>
            <Alert onClose={handleClose} severity="error">
              {errors}
            </Alert>
          </Snackbar>
        </Container>
      </Content>
    </Page>
  );
};

export default FurnitureDetail;
