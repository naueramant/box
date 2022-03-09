import { ButtonBase, Paper, styled } from "@mui/material";
import { FunctionComponent } from "react";

interface NavbarProps {
    children?: React.ReactNode | React.ReactNode[];
}

const NavbarContainer = styled(Paper)(({ theme }) => ({
    display: "flex",
    flexDirection: "column",
    padding: theme.spacing(2),
    borderRadius: 0,
    zIndex: theme.zIndex.appBar,
    borderRight: `1px solid ${theme.palette.divider}`,
}));

const Navbar: FunctionComponent<NavbarProps> = (props) => {
    return <NavbarContainer>{props.children}</NavbarContainer>;
};

interface NavbarItemProps {
    children?: React.ReactNode | React.ReactNode[];
    onClick?: () => void;
}

const NavbarItemContainer = styled(ButtonBase)(({ theme }) => ({
    paddingTop: theme.spacing(2),
    paddingBottom: theme.spacing(2),
    paddingLeft: theme.spacing(1),
    paddingRight: theme.spacing(1),
    display: "flex",
    flexDirection: "column",
    fontSize: theme.typography.fontSize,
    justifyContent: "center",
    alignItems: "center",
    cursor: "pointer",
    borderRadius: theme.shape.borderRadius,
    color: theme.palette.primary.main,
}));

const NavbarItem: FunctionComponent<NavbarItemProps> = (props) => {
    return <NavbarItemContainer onClick={props.onClick}>{props.children}</NavbarItemContainer>;
};

export { Navbar, NavbarItem };
