import React, { useCallback, useEffect, useState } from "react";

import TableRow from "@mui/material/TableRow";
import { styled } from "@mui/material/styles";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell, { tableCellClasses } from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";

import { ToastContainer, toast } from "react-toastify";
import "react-toastify/ReactToastify.min.css";
import { Button, IconButton, Paper, TableHead } from "@mui/material";
import { useDispatch, useSelector } from "react-redux";
import { FetchProductInHomePage } from "../../../../store/slices/ProductSlice";
import { useParams } from "react-router-dom";
import { setProductInBrand } from "../../../../store/slices/ProviderSlice";
import { currencyFormat } from "../../../../stogare_function/listActions";
import DeleteIcon from "@mui/icons-material/Delete";
import { ProductApi } from "../../../../api/ProductApi";

const StyledTableCell = styled(TableCell)(({ theme }) => ({
  [`&.${tableCellClasses.head}`]: {
    backgroundColor: theme.palette.common.black,
    color: theme.palette.common.white,
  },
  [`&.${tableCellClasses.body}`]: {
    fontSize: 14,
  },
}));

const StyledTableRow = styled(TableRow)(({ theme }) => ({
  // hide last border
  "&:last-child td, &:last-child th": {
    border: 0,
  },
}));

const noImgAvailable =
  "data:image/png;base64,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";

export const ListViewProducts = () => {
  const dispatch = useDispatch();
  let { id } = useParams();
  const filter = {
    "fields[]": `provider_id_${id}`,
    "sorts[]": "id_DESC",
  };

  const ListProduct = useSelector(
    (state) => state.product.ProductPreviewInHomePage
  );
  const listProductInBrand = useSelector(
    (state) => state.provider.ProductInBrand
  );

  const [isFirstRender, setIsFirstRender] = useState(true);
  const loadProduct = useCallback(async () => {
    await dispatch(FetchProductInHomePage(filter));
  });

  useEffect(() => {
    if (ListProduct.status != 200 && ListProduct.status != 204) {
      loadProduct();
    }
    if (ListProduct.status == 200) {
      if (isFirstRender) {
        setIsFirstRender(false);
        dispatch(setProductInBrand(ListProduct.data.data));
      }
    }
  }, [
    loadProduct,
    dispatch,
    ListProduct,
    listProductInBrand,
    id,
    isFirstRender,
  ]);

  const DeleteProductFromID = async(idProduct)=>{
    await ProductApi.DeleteProduct(idProduct)
    .then(res=>{
      if(res.status==200)
      {
        toast("Delete Item success", {
          type: "success",
          autoClose: 1000,
          onClose:setTimeout(()=>{
              window.location.reload()
          },1500)
        });
      }
    })
  }

  const handleDeleteButton = (e) =>{
    console.log(e.curren)
    DeleteProductFromID(e.currentTarget.id)

  }
  return (
    <div className="space-y-3">
      <h1 className="font-bold text-xl">Table product:</h1>
      <div>
        <ToastContainer position="top-right" newestOnTop />

        <TableContainer component={Paper}>
          <Table aria-label="customized table">
            <TableHead>
              <TableRow>
                <StyledTableCell>ID</StyledTableCell>

                <StyledTableCell>Image</StyledTableCell>
                <StyledTableCell align="center">Name Product</StyledTableCell>
                <StyledTableCell align="left">Price</StyledTableCell>
                <StyledTableCell align="center">Discount</StyledTableCell>
                <StyledTableCell align="center">Action</StyledTableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {listProductInBrand.length === 0 ? (
                <div></div>
              ) : (
                listProductInBrand.map((row) => (
                  <StyledTableRow key={row.id}>
                    <StyledTableCell align="left">{row.ID}</StyledTableCell>

                    <StyledTableCell
                      component="th"
                      scope="row"
                      sx={{ width: 150, padding: 1 }}
                    >
                      <img
                        src={
                          !row.Media ? noImgAvailable : row.Media[0].mediaPath
                        }
                        alt="Anh san pham"
                        className="w-[100px] h-[100px]"
                      ></img>
                    </StyledTableCell>
                    <StyledTableCell align="center">{row.Name}</StyledTableCell>

                    <StyledTableCell align="left">
                      {currencyFormat(row.Price)}đ
                    </StyledTableCell>
                    <StyledTableCell
                      align="center"
                      sx={{ width: 20, padding: 1 }}
                    >
                      <div className=" py-1 border border-[#C40201] text-[#C40201]">
                        {`${row.Discount}%`}
                      </div>
                    </StyledTableCell>
                    <StyledTableCell align="center">
                      <IconButton
                        id={row.ID}
                        color="primary"
                        aria-label="upload picture"
                        component="label"
                        onClick={handleDeleteButton}
                      >
                        <DeleteIcon />
                      </IconButton>
                    </StyledTableCell>
                  </StyledTableRow>
                ))
              )}
            </TableBody>
          </Table>
        </TableContainer>
      </div>
    </div>
  );
};
