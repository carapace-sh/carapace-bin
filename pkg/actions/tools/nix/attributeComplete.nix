all@{ __carapaceInput__, __carapaceAttrPath__, ... }:
let
  inputArgs = builtins.removeAttrs all ["__carapaceInput__" "__carapaceAttrPath__"];

  input =
    if builtins.isFunction __carapaceInput__ then
      __carapaceInput__ inputArgs
    else
      __carapaceInput__;

  autocall = fnOrAttrset:
    if builtins.isFunction fnOrAttrset then
      fnOrAttrset {}
    else
      fnOrAttrset;

  paths = builtins.filter (path: (builtins.isString path) && path != "") (builtins.split "\\." __carapaceAttrPath__);

  reducer = attrset: name: autocall (builtins.getAttr name attrset);
  result = builtins.foldl' reducer input paths;
in
  if builtins.isAttrs result then
    builtins.attrNames result
  else
    []
