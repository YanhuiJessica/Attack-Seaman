db.createUser(
    {
        user: "attackSeaman",
        pwd: "cuccs1sgreat",
        roles: [
            {
                role: "readWrite",
                db: "attackSeaman"
            }
        ]
    }
);