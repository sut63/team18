import React, { FC, useEffect } from 'react';
import { Theme, makeStyles, withStyles, createStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import Typography from '@material-ui/core/Typography';
import Divider from '@material-ui/core/Divider';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow'
import Paper from '@material-ui/core/Paper';
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
import { EntDataRoom, EntFurnitureDetail } from '../../api';
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
    bottom: {
        marginLeft: theme.spacing(3),
        marginTop: theme.spacing(1),
        marginBottom: theme.spacing(2),
    },
    divider: {
        margin: theme.spacing(2, 0),
    },
    table: {
        minWidth: 650,
    },
}));


const SearchFurniture: FC<{}> = () => {
    const classes = useStyles();
    const api = new DefaultApi();

    // ดึงคุกกี้
    var ck = new Cookies()
    var cookieName = ck.GetCookie()

    // Dataroom
    const [idDataroom, setIdDataroom] = React.useState<number>(0)
    const [dataroom, setDataroom] = React.useState<EntDataRoom[]>([])
    const getDataroom = async () => {
        const res = await api.listDataroom({})
        setDataroom(res)
    }

    // alert setting
    const [open, setOpen] = React.useState(false);
    const [fail, setFail] = React.useState(false);

    //close alert 
    const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
        if (reason === 'clickaway') {
            return;
        }
        setFail(false);
        setOpen(false);
    };


    // CheckIn  
    var lenfurnituredetail : number
    const [furnituredetail, setFurnitureDetail] = React.useState<EntFurnitureDetail[]>([])
    const getFurnitureDetail = async () => {
        const res = await api.listFurnitureDetailRoom({ id: idDataroom })
        setFurnitureDetail(res)
        lenfurnituredetail = res.length
        if (lenfurnituredetail > 0){
            setOpen(true)
        }else{
            setFail(true)
        }   
    }

    // Lifecycle Hooks
    useEffect(() => {
        getDataroom();
    }, []);
    

    // set data to object and setIdcustomer
    const handleChange = (
        event: React.ChangeEvent<{ name?: string; value: any }>,
    ) => {
        const { value } = event.target;
        setIdDataroom(value);
    };

    // clear cookies
    function Clears() {
        ck.ClearCookie()
        window.location.reload(false)
    }

    // function seach data
    function seach() {
        getFurnitureDetail();
    }

    return (
        <Page theme={pageTheme.home}>
            <Header style={HeaderCustom} title={`ระบบค้นหาเฟอร์นิเจอร์สำหรับห้องพัก`}>
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
                <Grid container spacing={1}>
                    <Grid item xs={1}>
                        <div className={classes.paper}><h3>ห้องพัก</h3></div>
                    </Grid>
                    <Grid item xs={3}>
                        <FormControl variant="outlined" className={classes.formControl}>
                            <InputLabel>เลือกห้องพัก</InputLabel>
                            <Select
                                name="Dataroom"
                                value={idDataroom || ''} // (undefined || '') = ''
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
                    <Grid item xs={6}>
                        <Button
                            variant="contained"
                            color="secondary"
                            size="large"
                            onClick={seach}
                        >
                            search
                        </Button>
                    </Grid>
                </Grid>

                <Grid item xs={12}>
                    <Divider className={classes.divider} />
                    <Typography variant="subtitle1" gutterBottom>
                        ตารางเฟอร์นิเจอร์:
                        </Typography>
                </Grid>
                <TableContainer component={Paper}>
                    <Table className={classes.table} aria-label="simple table">
                        <TableHead>
                            <TableRow>
                                <TableCell align="center">ห้องพัก</TableCell>
                                <TableCell align="center">เฟอร์นิเจอร์</TableCell>
                                <TableCell align="center">พนักงาน</TableCell>
                                <TableCell align="center">รายละเอียด</TableCell>
                                <TableCell align="center">สีของเฟอร์นิเจอร์</TableCell>
                                <TableCell align="center">จำนวน</TableCell>
                                <TableCell align="center">เวลา</TableCell>
                                
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {furnituredetail.map((item: EntFurnitureDetail) => (
                                <TableRow key={item.id}>
                                    <TableCell align="center">{item.edges?.rooms?.roomnumber}</TableCell>
                                    <TableCell align="center">{item.edges?.furnitures?.furnitureName}</TableCell>
                                    <TableCell align="center">{item.edges?.counterstaffs?.name}</TableCell>                                   
                                    <TableCell align="center">{item.detail}</TableCell>
                                    <TableCell align="center">{item.furnitureColour}</TableCell>
                                    <TableCell align="center">{item.furnitureAmount}</TableCell>
                                    <TableCell align="center">{item.dateAdd}</TableCell>                                   
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>
                <Snackbar open={open} autoHideDuration={6000} onClose={handleClose}>
            <Alert onClose={handleClose} severity="success">
              ค้นหาสำเร็จ
          </Alert>
          </Snackbar>

          <Snackbar open={fail} autoHideDuration={6000} onClose={handleClose}>
            <Alert onClose={handleClose} severity="error">
              ไม่พบข้อมูล
          </Alert>
          </Snackbar>
            </Content>
        </Page>
    );
};

export default SearchFurniture;
