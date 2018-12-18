// Copyright 2018 tinystack Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package project

import (
    "time"

    "github.com/tinystack/syncd/model"
)

func Create(data *Project) bool {
    data.Utime = int(time.Now().Unix())
    return model.Create(TableName, data)
}

func List(query model.QueryParam) ([]Project, bool) {
    var p []Project
    ok := model.GetMulti(TableName, &p, query)
    return p, ok
}

func Total(query model.QueryParam) (int, bool) {
    var count int
    ok := model.Count(TableName, &count, query)
    return count, ok
}

func Get(id int) (Project, bool){
    var data Project
    ok := model.GetByPk(TableName, &data, id)
    return data, ok
}

func Update(id int, data Project) bool {
    updateFields := map[string]interface{}{
        "name": data.Name,
        "description": data.Description,
        "space": data.Space,
        "repo": data.Repo,
        "repo_url": data.RepoUrl,
        "deploy_server": data.DeployServer,
        "deploy_user": data.DeployUser,
        "deploy_path": data.DeployPath,
        "deploy_history": data.DeployHistory,
        "pre_deploy_cmd": data.PreDeployCmd,
        "post_deploy_cmd": data.PostDeployCmd,
        "need_audit": data.NeedAudit,
        "status": data.Status,
        "repo_user": data.RepoUser,
        "repo_pass": data.RepoPass,
        "repo_mode": data.RepoMode,
        "build_script": data.BuildScript,
        "utime": int(time.Now().Unix()),
    }
    ok := model.Update(TableName, updateFields, model.QueryParam{
        Where: []model.WhereParam{
            model.WhereParam{
                Field: "id",
                Prepare: id,
            },
        },
    })
    return ok
}

func Delete(id int) bool {
    ok := model.DeleteByPk(TableName, Project{ID: id})
    return ok
}