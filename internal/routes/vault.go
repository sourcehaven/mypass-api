package routes

//func (h *Handler) GetVaults(c *fiber.Ctx) error {
//	// You can access the user claims like this, if needed:
//	user := c.Locals("user").(*jwt.Token)
//	claims := user.Claims.(jwt.MapClaims)
//	userId := int(claims["sub"].(float64)) // The user ID from 'sub' claim
//
//	vaults, err := h.Store.GetVaults(userId)
//
//	if err != nil {
//		SendResponse(c, http.StatusBadRequest, "", nil)
//		return err
//	}
//
//	// Perform the authorization check for the user based on the claims
//	// and the operation they're trying to perform.
//
//	SendResponse(c, http.StatusOK, "", vaults)
//	return nil
//}
