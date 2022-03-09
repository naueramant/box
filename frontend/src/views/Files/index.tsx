import {
    AccessTimeOutlined,
    FolderOutlined,
    GridViewOutlined,
    PersonOutline,
    Search,
    SearchOutlined,
    ShareOutlined,
    TableRowsOutlined,
} from "@mui/icons-material";
import {
    Avatar,
    Divider,
    IconButton,
    InputBase,
    Paper,
    Stack,
    styled,
    ToggleButton,
    ToggleButtonGroup,
    Typography,
} from "@mui/material";
import { Navbar, NavbarItem } from "components/Navbar";
import { FlexSpacer } from "components/Spacer";
import { stringAvatar } from "helpers/avatar";
import { useUser } from "hooks/user";
import { FunctionComponent } from "react";
import { useNavigate } from "react-router-dom";

interface FilesViewProps {}

const Container = styled("div")(({ theme }) => ({
    height: "100%",
    display: "flex",
    flexDirection: "row",
}));

const FilesContainer = styled(Paper)(({ theme }) => ({
    flex: 1,
    padding: theme.spacing(4),
    backgroundColor: theme.palette.background.default,
}));

const SidebarContainer = styled(Paper)(({ theme }) => ({
    width: "300px",
    padding: theme.spacing(4),
}));

const StyledDivider = styled(Divider)(({ theme }) => ({
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
}));

const FilesView: FunctionComponent<FilesViewProps> = () => {
    const navigate = useNavigate();
    const [user] = useUser();

    return (
        <Container>
            <Navbar>
                <NavbarItem>
                    <FolderOutlined fontSize="large" />
                    All
                </NavbarItem>
                <NavbarItem>
                    <SearchOutlined />
                    Search
                </NavbarItem>
                <NavbarItem>
                    <AccessTimeOutlined />
                    Recent
                </NavbarItem>
                <NavbarItem>
                    <ShareOutlined />
                    Shared
                </NavbarItem>

                <FlexSpacer />

                {user && (
                    <NavbarItem onClick={() => navigate("/login")}>
                        <Avatar variant="rounded" {...stringAvatar(user.username)} />
                    </NavbarItem>
                )}
            </Navbar>

            <FilesContainer>
                <Stack direction="row" alignItems="center">
                    <Typography variant="h5" fontWeight={600}>
                        Everything
                    </Typography>

                    <FlexSpacer />

                    <Paper
                        variant="outlined"
                        sx={{ p: "2px 4px", display: "flex", alignItems: "center", width: 400, marginRight: 2 }}
                    >
                        <InputBase sx={{ ml: 1, flex: 1 }} placeholder="Search" />
                        <IconButton type="submit" sx={{ p: "10px" }} aria-label="search">
                            <Search />
                        </IconButton>
                    </Paper>

                    <ToggleButtonGroup>
                        <ToggleButton value="list">
                            <TableRowsOutlined />
                        </ToggleButton>
                        <ToggleButton value="grid">
                            <GridViewOutlined />
                        </ToggleButton>
                    </ToggleButtonGroup>
                </Stack>
                <StyledDivider />
            </FilesContainer>
            {/* 
            <SidebarContainer>
                <Typography variant="h5">File Preview</Typography>
            </SidebarContainer> */}
        </Container>
    );
};

export default FilesView;
