/* Do not change, this code is generated from Golang structs */


export interface ApiGetVersionResponse {
    app: string;
}
export interface ApiWorkspaceTreeEntry {
    path: string;
    title: string;
    type: string;
    children?: ApiWorkspaceTreeEntry[];
}
export interface ApiGetWorkspaceResponse {
    description: string;
    variables: {[key: string]: string};
    tree: ApiWorkspaceTreeEntry[];
}
export interface ApiServerError {
    error: string;
}