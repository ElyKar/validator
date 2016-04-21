### validator ###

Motivation
==========

Validation of data is necessary when woking with computers,
because when working with input data we have to ensure its cohesion.
It can be used to validate a forms, or simply for reading config files.

Making validaters atomic is a way to reuse them accross the code, and it's
very frequent that a validater checks only for one condition

Functionality
=============

- Exec : Runs a serie of Validaters and returns the first error found
- Collect : Runs a serie of Validaters and returns all the errors found along the way
- Using the ValidaterSet, it is possible to use make a hierarchy between validaters : collect errors for the first group, and if no errors have been raised then continue the processing.

First Example
=============

Given this structure :

    type Config struct {
        Age int
        Name string
        Pet string
    }

It is possible to validate it in every way imaginable.
If we want to collect all errors, then calling the following will do the job :

    err := Collect(newIntGreaterThanValidater(config.Age, 0), newIntLowerThan(config.Age, 100), newStringNotEmptyValidater(config.Name), newStringNotEmptyValidater(config.Pet), newHasATurtlePetValidater(config))

Second Example
==============

Now, maybe some validation are necessary to the others.
Let's say that we want to check if the content of two files is equal, this is were ValidaterSet becomes handy :
We want to check content of fileA against content of fileB.
One might write the following :

    filesExists := NewValidaterSet(newDoesFileExist("fileA"), newDoesFileExists("fileB"))
    filesAreReadable := NewValidaterSet(newIsFileReadable("fileA"), newIsFileReadable("fileB"))
    err := Exec(filesExists, filesAreReadable, newAreFileContentEquals("fileA", "fileB"))
    if err != nil {
        return err
    }

Conclusion
==========

There are few functionalities yet they are very flexible and can adapt to multiple situations.
It's up to you now to create the adequate validater chain !
