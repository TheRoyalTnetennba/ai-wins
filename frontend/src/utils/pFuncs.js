export const sameArr = (rc1, rc2) => {
  if (rc1 === rc2) return true;
  if (rc1 == null || rc2 == null) return false;
  if (rc1.length != rc2.length) return false;
  for (let i = 0; i < rc1.length; ++i) {
    if (rc1[i] !== rc2[i]) return false;
  }
  return true;
}
