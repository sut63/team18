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
import { EntCustomer, EntReserveRoom } from '../../api';
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


const SearchReserveRoom: FC<{}> = () => {
    const classes = useStyles();
    const api = new DefaultApi();

    // ดึงคุกกี้
    var ck = new Cookies()
    var cookieName = ck.GetCookie()
    var cookieID = ck.GetID()

    // Customer
    const [idCustomer, setIdcustomer] = React.useState<number>(0)
    const [customer, setCustomer] = React.useState<EntCustomer[]>([])
    const getCustomer = async () => {
        const res = await api.listCustomer({})
        setCustomer(res)
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
    var checkReserve: number
    // ReserveRoom  
    const [reserveroom, setReserveroom] = React.useState<EntReserveRoom[]>([])
    const getReseveroom = async () => {
        const res = await api.listReserveCustomer({ id: idCustomer })
        setReserveroom(res)
        checkReserve = res.length
        if (checkReserve > 0) {
            setOpen(true)
        } else {
            setFail(true)
        }
    }

    // Lifecycle Hooks
    useEffect(() => {
        getCustomer();
    }, []);

    // set data to search reserveroom
    const handleChange = (
        event: React.ChangeEvent<{ value: any }>,
    ) => {
        const { value } = event.target;
        setIdcustomer(value);
    };

    // clear cookies
    function Clears() {
        ck.ClearCookie()
        window.location.reload(false)
    }

    // function seach data
    function seach() {
        getReseveroom();
    }

    return (
        <Page theme={pageTheme.home}>
            <Header style={HeaderCustom} title={`ระบบค้นหาใบจอง`}>
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
                        <div className={classes.paper}>customer</div>
                    </Grid>
                    <Grid item xs={3}>
                        <FormControl variant="outlined" className={classes.formControl}>
                            <InputLabel>select customer</InputLabel>
                            <Select
                                name="Customer"
                                value={idCustomer || ''} // (undefined || '') = ''
                                onChange={handleChange}
                            >
                                {customer.map(item => {
                                    return (
                                        <MenuItem key={item.id} value={item.id}>
                                            {item.name}
                                        </MenuItem>
                                    );
                                })}
                            </Select>
                        </FormControl>
                    </Grid>
                    <Grid item xs={3}>
                        <Button
                            className={classes.bottom}
                            variant="contained"
                            color="primary"
                            size="large"
                            onClick={seach}
                        >
                            ค้นหา
                        </Button>
                    </Grid>
                </Grid>

                <Grid item xs={12}>
                    <Divider className={classes.divider} />
                    <Typography variant="subtitle1" gutterBottom>
                        ตารางใบจอง:
                        </Typography>
                </Grid>
                <TableContainer component={Paper}>
                    <Table className={classes.table} aria-label="simple table">
                        <TableHead>
                            <TableRow>
                                <TableCell align="center">เลขใบจอง</TableCell>
                                <TableCell align="center">ห้องที่จอง</TableCell>
                                <TableCell align="center">สถานะ</TableCell>
                                <TableCell align="center">จำนวนคนที่เข้าพัก</TableCell>
                                <TableCell align="center">เบอร์โทรติดต่อ</TableCell>
                                <TableCell align="center">วันที่เข้าพัก</TableCell>
                                <TableCell align="center">ราคาสุทธิ</TableCell>
                                <TableCell align="center">request</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {reserveroom.map((item: EntReserveRoom) => (
                                <TableRow key={item.id}>
                                    <TableCell align="center">{item.id}</TableCell>
                                    <TableCell align="center">{item.edges?.room?.roomnumber}</TableCell>
                                    <TableCell align="center">{item.edges?.status?.statusName}</TableCell>
                                    <TableCell align="center">{item.amount}</TableCell>
                                    <TableCell align="center">{item.phoneNumber}</TableCell>
                                    <TableCell align="center">{item.reserveDate}</TableCell>
                                    <TableCell align="center">{item.netPrice} บาท</TableCell>
                                    <TableCell align="center">{item.request}</TableCell>
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

export default SearchReserveRoom;
