if d.Get("deletion_protection").(bool) {
	return fmt.Errorf("cannot destroy database without setting deletion_protection=false and running `terraform apply`")
}
